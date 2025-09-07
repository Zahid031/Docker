package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task-service/rabbitmq"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	UserID      uint      `json:"user_id" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var db *gorm.DB
var taskPublisher *rabbitmq.Publisher
var consumer *rabbitmq.Consumer

func main() {
	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgresql://postgres:postgres123@postgres:5432/todo_db?sslmode=disable"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&Task{})

	// Setup RabbitMQ
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		rabbitmqURL = "amqp://admin:admin123@rabbitmq:5672/" // Use service name
	}

	// Wait for RabbitMQ to be ready
	waitForRabbitMQ(rabbitmqURL)

	// Initialize RabbitMQ publisher
	taskPublisher, err = rabbitmq.NewPublisher(rabbitmqURL)
	if err != nil {
		log.Printf("Failed to initialize RabbitMQ publisher: %v", err)
	} else {
		log.Println("RabbitMQ publisher initialized successfully")
	}

	// Initialize and start RabbitMQ consumer
	consumer, err = rabbitmq.NewConsumer(rabbitmqURL, db)
	if err != nil {
		log.Printf("Failed to initialize RabbitMQ consumer: %v", err)
	} else {
		go func() {
			err = consumer.StartConsuming()
			if err != nil {
				log.Printf("RabbitMQ consumer error: %v", err)
			}
		}()
		log.Println("RabbitMQ consumer started successfully")
	}

	// Setup graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Setup Gin router
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Routes
	api := r.Group("/api/tasks")
	{
		api.GET("/", getTasks)
		api.POST("/", createTask)
		api.GET("/:id", getTask)
		api.PUT("/:id", updateTask)
		api.DELETE("/:id", deleteTask)
		api.GET("/user/:user_id", getTasksByUser)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// RabbitMQ health check
	r.GET("/health/rabbitmq", func(c *gin.Context) {
		if taskPublisher != nil {
			c.JSON(http.StatusOK, gin.H{"rabbitmq": "healthy"})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{"rabbitmq": "unhealthy"})
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Task service starting on port %s", port)
		if err := r.Run(":" + port); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	<-c
	log.Println("Shutting down...")

	// Cleanup
	if consumer != nil {
		consumer.Close()
	}
	if taskPublisher != nil {
		taskPublisher.Close()
	}

	log.Println("Shutdown complete")
}

func waitForRabbitMQ(rabbitmqURL string) {
	maxRetries := 30
	for i := 0; i < maxRetries; i++ {
		conn, err := amqp.Dial(rabbitmqURL)
		if err == nil {
			conn.Close()
			log.Println("RabbitMQ is ready")
			return
		}
		log.Printf("Waiting for RabbitMQ... (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}
	log.Println("RabbitMQ connection timeout, continuing anyway...")
}

func getTasks(c *gin.Context) {
	var tasks []Task
	result := db.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Publish task created event
	if taskPublisher != nil {
		taskEvent := rabbitmq.TaskEvent{
			TaskID:      task.ID,
			Title:       task.Title,
			Description: task.Description,
			UserID:      task.UserID,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt,
		}
		
		err := taskPublisher.PublishTaskEvent("task.created", taskEvent)
		if err != nil {
			log.Printf("Failed to publish task.created event: %v", err)
		} else {
			log.Printf("Published task.created event for task %d", task.ID)
		}
	}

	c.JSON(http.StatusCreated, task)
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if result := db.First(&task, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if result := db.First(&task, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	var updateData Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task
	if result := db.Model(&task).Updates(updateData); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Refresh the task data from database to get updated timestamp
	db.First(&task, id)

	// Publish task updated event
	if taskPublisher != nil {
		taskEvent := rabbitmq.TaskEvent{
			TaskID:      task.ID,
			Title:       task.Title,
			Description: task.Description,
			UserID:      task.UserID,
			Completed:   task.Completed,
			UpdatedAt:   task.UpdatedAt,
		}
		
		err := taskPublisher.PublishTaskEvent("task.updated", taskEvent)
		if err != nil {
			log.Printf("Failed to publish task.updated event: %v", err)
		} else {
			log.Printf("Published task.updated event for task %d", task.ID)
		}
	}

	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if result := db.First(&task, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Store task info before deletion for event
	taskEvent := rabbitmq.TaskEvent{
		TaskID:      task.ID,
		Title:       task.Title,
		Description: task.Description,
		UserID:      task.UserID,
		Completed:   task.Completed,
		DeletedAt:   time.Now(),
	}

	// Delete the task
	if result := db.Delete(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Publish task deleted event
	if taskPublisher != nil {
		err := taskPublisher.PublishTaskEvent("task.deleted", taskEvent)
		if err != nil {
			log.Printf("Failed to publish task.deleted event: %v", err)
		} else {
			log.Printf("Published task.deleted event for task %d", task.ID)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func getTasksByUser(c *gin.Context) {
	userID := c.Param("user_id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var tasks []Task
	result := db.Where("user_id = ?", userIDInt).Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}