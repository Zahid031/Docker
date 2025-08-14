package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type UserEvent struct {
	EventType string `json:"event_type"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Timestamp string `json:"timestamp"`
}

type KafkaConsumer struct {
	reader *kafka.Reader
	db     *gorm.DB
}

func NewKafkaConsumer(db *gorm.DB) *KafkaConsumer {
	brokers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokers},
		Topic:   "user-events",
		GroupID: "task-service-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &KafkaConsumer{
		reader: reader,
		db:     db,
	}
}

func (kc *KafkaConsumer) StartConsumer(ctx context.Context) {
	log.Println("Starting Kafka consumer for user events...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping Kafka consumer...")
			kc.reader.Close()
			return
		default:
			message, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading Kafka message: %v", err)
				continue
			}

			var event UserEvent
			if err := json.Unmarshal(message.Value, &event); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			log.Printf("Received event: %+v", event)

			switch event.EventType {
			case "user_created":
				kc.handleUserCreated(event)
			case "user_deleted":
				kc.handleUserDeleted(event)
			default:
				log.Printf("Unknown event type: %s", event.EventType)
			}
		}
	}
}

func (kc *KafkaConsumer) handleUserCreated(event UserEvent) {
	log.Printf("Handling user created event for user %d", event.UserID)

	// Create default tasks for new user
	defaultTasks := []Task{
		{
			Title:       "Welcome to Todo App!",
			Description: "This is your first task. Click to mark it complete!",
			Completed:   false,
			UserID:      event.UserID,
		},
		{
			Title:       "Explore the features",
			Description: "Try creating, updating, and deleting tasks",
			Completed:   false,
			UserID:      event.UserID,
		},
		{
			Title:       "Set up your profile",
			Description: "Complete your profile information",
			Completed:   false,
			UserID:      event.UserID,
		},
	}

	for _, task := range defaultTasks {
		if result := kc.db.Create(&task); result.Error != nil {
			log.Printf("Error creating default task for user %d: %v", event.UserID, result.Error)
		} else {
			log.Printf("Created default task '%s' for user %d", task.Title, event.UserID)
		}
	}
}

func (kc *KafkaConsumer) handleUserDeleted(event UserEvent) {
	log.Printf("Handling user deleted event for user %d", event.UserID)

	// Delete all tasks for the deleted user
	result := kc.db.Where("user_id = ?", event.UserID).Delete(&Task{})
	if result.Error != nil {
		log.Printf("Error deleting tasks for user %d: %v", event.UserID, result.Error)
	} else {
		log.Printf("Deleted %d tasks for user %d", result.RowsAffected, event.UserID)
	}
}

func (kc *KafkaConsumer) Close() {
	if kc.reader != nil {
		kc.reader.Close()
	}
}