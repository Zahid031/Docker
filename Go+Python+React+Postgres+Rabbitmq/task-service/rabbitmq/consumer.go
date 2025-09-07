package rabbitmq

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type UserEvent struct {
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Consumer struct {
	rabbitmqURL string
	db          *gorm.DB
	conn        *amqp.Connection
	channel     *amqp.Channel
	done        chan bool
	reconnect   chan bool
}

func NewConsumer(rabbitmqURL string, db *gorm.DB) (*Consumer, error) {
	consumer := &Consumer{
		rabbitmqURL: rabbitmqURL,
		db:          db,
		done:        make(chan bool),
		reconnect:   make(chan bool),
	}

	err := consumer.connect()
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (c *Consumer) connect() error {
	var err error
	
	// Configure connection with heartbeat settings
	config := amqp.Config{
		Heartbeat: 30 * time.Second,  // Reduced heartbeat interval
		Locale:    "en_US",
		Dial: amqp.DefaultDial(30 * time.Second), // Connection timeout
	}

	c.conn, err = amqp.DialConfig(c.rabbitmqURL, config)
	if err != nil {
		return err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		c.conn.Close()
		return err
	}

	// Set QoS to limit unacknowledged messages
	err = c.channel.Qos(
		10,    // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		c.channel.Close()
		c.conn.Close()
		return err
	}

	// Declare the queue with proper settings
	_, err = c.channel.QueueDeclare(
		"task_service_queue", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		amqp.Table{
			"x-message-ttl": 300000, // 5 minutes TTL
		},
	)
	if err != nil {
		c.channel.Close()
		c.conn.Close()
		return err
	}

	// Listen for connection close events
	go c.watchConnection()

	log.Println("Connected to RabbitMQ successfully")
	return nil
}

func (c *Consumer) watchConnection() {
	notifyConnClose := make(chan *amqp.Error)
	notifyChannelClose := make(chan *amqp.Error)
	
	c.conn.NotifyClose(notifyConnClose)
	c.channel.NotifyClose(notifyChannelClose)

	select {
	case err := <-notifyConnClose:
		if err != nil {
			log.Printf("Connection closed: %v", err)
		}
		c.reconnect <- true
	case err := <-notifyChannelClose:
		if err != nil {
			log.Printf("Channel closed: %v", err)
		}
		c.reconnect <- true
	case <-c.done:
		return
	}
}

func (c *Consumer) StartConsuming() error {
	for {
		select {
		case <-c.done:
			log.Println("Consumer stopped")
			return nil
		case <-c.reconnect:
			log.Println("Attempting to reconnect to RabbitMQ...")
			c.cleanup()
			
			// Wait before reconnecting
			time.Sleep(5 * time.Second)
			
			err := c.connect()
			if err != nil {
				log.Printf("Failed to reconnect: %v", err)
				time.Sleep(10 * time.Second)
				c.reconnect <- true // Try again
				continue
			}
			
			// Start consuming again
			go c.consume()
		default:
			// Start initial consumption
			go c.consume()
			
			// Block until we need to reconnect or stop
			select {
			case <-c.reconnect:
				continue
			case <-c.done:
				return nil
			}
		}
	}
}

func (c *Consumer) consume() {
	msgs, err := c.channel.Consume(
		"task_service_queue", // queue
		"task-consumer",      // consumer tag
		false,                // auto-ack (set to false for manual ack)
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Printf("Failed to register consumer: %v", err)
		c.reconnect <- true
		return
	}

	log.Println("Consumer started, waiting for messages...")

	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				log.Println("Message channel closed")
				c.reconnect <- true
				return
			}
			
			// Process message with timeout
			go func(delivery amqp.Delivery) {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Panic in message handler: %v", r)
						delivery.Nack(false, true) // Requeue on panic
					}
				}()
				
				err := c.handleUserEvent(delivery.RoutingKey, delivery.Body)
				if err != nil {
					log.Printf("Failed to handle message: %v", err)
					delivery.Nack(false, true) // Requeue on error
				} else {
					delivery.Ack(false) // Acknowledge successful processing
				}
			}(d)
			
		case <-c.done:
			return
		}
	}
}

func (c *Consumer) handleUserEvent(routingKey string, body []byte) error {
	log.Printf("Processing message: %s - %s", routingKey, string(body))

	var event UserEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return err
	}

	switch routingKey {
	case "user.created":
		return c.handleUserCreated(event)
	case "user.updated":
		return c.handleUserUpdated(event)
	case "user.deleted":
		return c.handleUserDeleted(event)
	default:
		log.Printf("Unknown routing key: %s", routingKey)
		return nil
	}
}

func (c *Consumer) handleUserCreated(event UserEvent) error {
	log.Printf("User created: %d - %s", event.UserID, event.Name)
	
	welcomeTask := Task{
		Title:       "Welcome to Todo App!",
		Description: "This is your first task. Start organizing your life!",
		UserID:      event.UserID,
		Completed:   false,
	}
	
	if err := c.db.Create(&welcomeTask).Error; err != nil {
		log.Printf("Failed to create welcome task for user %d: %v", event.UserID, err)
		return err
	}
	
	log.Printf("Created welcome task for user %d", event.UserID)
	return nil
}

func (c *Consumer) handleUserUpdated(event UserEvent) error {
	log.Printf("User updated: %d - %s", event.UserID, event.Name)
	return nil
}

func (c *Consumer) handleUserDeleted(event UserEvent) error {
	log.Printf("User deleted: %d", event.UserID)
	
	result := c.db.Where("user_id = ?", event.UserID).Delete(&Task{})
	if result.Error != nil {
		log.Printf("Failed to delete tasks for user %d: %v", event.UserID, result.Error)
		return result.Error
	}
	
	log.Printf("Deleted %d tasks for user %d", result.RowsAffected, event.UserID)
	return nil
}

func (c *Consumer) cleanup() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Consumer) Close() {
	close(c.done)
	c.cleanup()
}