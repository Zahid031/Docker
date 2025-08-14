# Todo Microservices Application

A modern, event-driven microservices application built with Django REST Framework, Go (Gin), Apache Kafka (KRaft), and React. This application demonstrates real-time task management with automatic task creation and cleanup through Kafka events.

## 🏗️ Architecture Overview

```
Frontend (React) ↔ User Service (Django) ↔ Kafka ↔ Task Service (Go)
                          ↓                           ↓
                    PostgreSQL ←――――――――――――――→ PostgreSQL
```

### Services:
- **User Service**: Django REST Framework (Port 8000) - Manages user operations
- **Task Service**: Go with Gin (Port 8001) - Manages task operations  
- **Frontend**: React Application (Port 3000) - User interface
- **Kafka**: Event streaming platform (Port 9092) - Handles inter-service communication
- **PostgreSQL**: Database (Port 5432) - Data persistence

## ✨ Key Features

### Event-Driven Architecture
- **User Creation**: Automatically creates 3 welcome tasks for new users
- **User Deletion**: Automatically deletes all associated tasks when a user is deleted
- **Real-time Communication**: Services communicate asynchronously through Kafka events

### API Capabilities
- Full CRUD operations for users and tasks
- RESTful API design with proper HTTP status codes
- CORS enabled for frontend integration
- Comprehensive error handling and validation

## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Git

### 1. Clone and Setup
```bash
git clone <your-repository>
cd todo-microservices
```

### 2. Start the Application
```bash
# Build and start all services
docker-compose up --build

# Or run in background
docker-compose up --build -d
```

### 3. Verify Services
```bash
# Check all services are running
docker-compose ps

# Check service health
curl http://localhost:8000/api/users/     # User Service
curl http://localhost:8001/health         # Task Service  
curl http://localhost:3000               # Frontend
```

## 📊 Service Status

Once started, you can access:
- **Frontend Application**: http://localhost:3000
- **User Service API**: http://localhost:8000
- **Task Service API**: http://localhost:8001
- **Database**: localhost:5432 (postgres/postgres123)
- **Kafka**: localhost:9092

## 🔄 Event-Driven Workflow Demo

### Test the Complete Workflow:

```bash
# 1. Create a user (triggers automatic task creation)
curl -X POST http://localhost:8000/api/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe", 
    "email": "john@example.com"
  }'

# Response: User created with ID 1

# 2. Check automatically created welcome tasks
curl http://localhost:8001/api/tasks/user/1

# Response: 3 default tasks automatically created:
# - "Welcome to Todo App!"
# - "Explore the features" 
# - "Set up your profile"

# 3. Create additional tasks manually
curl -X POST http://localhost:8001/api/tasks/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries",
    "description": "Get milk and bread", 
    "user_id": 1
  }'

# 4. View all user tasks
curl http://localhost:8001/api/tasks/user/1

# Response: Shows 4 tasks (3 auto-created + 1 manual)

# 5. Delete the user (triggers automatic task cleanup)
curl -X DELETE http://localhost:8000/api/users/1/

# 6. Verify all tasks were automatically deleted
curl http://localhost:8001/api/tasks/user/1

# Response: Empty array - all tasks automatically cleaned up
```

## 📚 API Documentation

### User Service API (Port 8000)

#### User Model
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com", 
  "created_at": "2024-12-19T10:30:00Z",
  "updated_at": "2024-12-19T10:30:00Z"
}
```

#### Endpoints

| Method | Endpoint | Description | Events Triggered |
|--------|----------|-------------|------------------|
| GET | `/api/users/` | Get all users | None |
| POST | `/api/users/` | Create user | ✅ `user_created` → Auto-creates 3 tasks |
| GET | `/api/users/{id}/` | Get specific user | None |
| PUT | `/api/users/{id}/` | Update user | None |
| DELETE | `/api/users/{id}/` | Delete user | ✅ `user_deleted` → Deletes all user tasks |

#### Create User Example
```bash
curl -X POST http://localhost:8000/api/users/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com"
  }'
```

### Task Service API (Port 8001)

#### Task Model
```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Get milk, bread, and eggs",
  "completed": false,
  "user_id": 1,
  "created_at": "2024-12-19T10:30:00Z", 
  "updated_at": "2024-12-19T10:30:00Z"
}
```

#### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tasks/` | Get all tasks |
| POST | `/api/tasks/` | Create task |
| GET | `/api/tasks/{id}` | Get specific task |
| PUT | `/api/tasks/{id}` | Update task |
| DELETE | `/api/tasks/{id}` | Delete task |
| GET | `/api/tasks/user/{user_id}` | Get user's tasks |
| GET | `/health` | Health check |

#### Create Task Example
```bash
curl -X POST http://localhost:8001/api/tasks/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete project",
    "description": "Finish the microservices app",
    "user_id": 1,
    "completed": false
  }'
```

## 🎯 Kafka Events

### Event Types

#### 1. User Created Event
**Topic**: `user-events`
**Trigger**: When a new user is created
**Action**: Task service automatically creates 3 welcome tasks

```json
{
  "event_type": "user_created",
  "user_id": 1,
  "user_name": "John Doe", 
  "user_email": "john@example.com",
  "timestamp": "2024-12-19T10:30:00Z"
}
```

#### 2. User Deleted Event  
**Topic**: `user-events`
**Trigger**: When a user is deleted
**Action**: Task service automatically deletes all user's tasks

```json
{
  "event_type": "user_deleted",
  "user_id": 1,
  "timestamp": "2024-12-19T11:00:00Z"
}
```

## 🛠️ Development & Monitoring

### View Service Logs
```bash
# All services
docker-compose logs -f

# Specific services
docker-compose logs -f todo_user_service
docker-compose logs -f todo_task_service  
docker-compose logs -f todo_kafka
docker-compose logs -f todo_frontend
```

### Monitor Kafka Events
```bash
# Access Kafka container
docker exec -it todo_kafka bash

# View all topics
kafka-topics --bootstrap-server localhost:9092 --list

# Monitor user events in real-time
kafka-console-consumer --bootstrap-server localhost:9092 \
  --topic user-events --from-beginning
```

### Database Access
```bash
# Connect to PostgreSQL
docker exec -it todo_postgres psql -U postgres -d todo_db

# View users table
SELECT * FROM users_user;

# View tasks table  
SELECT * FROM tasks;
```

## 🔧 Configuration

### Environment Variables

#### User Service
- `DEBUG=1` - Django debug mode
- `DATABASE_URL` - PostgreSQL connection string
- `ALLOWED_HOSTS` - Django allowed hosts
- `CORS_ALLOWED_ORIGINS` - CORS origins
- `KAFKA_BOOTSTRAP_SERVERS` - Kafka connection

#### Task Service  
- `DATABASE_URL` - PostgreSQL connection string
- `PORT=8001` - Service port
- `USER_SERVICE_URL` - User service URL
- `KAFKA_BOOTSTRAP_SERVERS` - Kafka connection

## 🚨 Troubleshooting

### Common Issues

#### Services won't start
```bash
# Check service status
docker-compose ps

# Restart specific service
docker-compose restart todo_user_service
```

#### Kafka connection issues
```bash
# Check Kafka health
docker exec -it todo_kafka kafka-topics --bootstrap-server localhost:9092 --list

# Restart Kafka
docker-compose restart todo_kafka
```

#### Database connection issues
```bash
# Check database health
docker exec -it todo_postgres pg_isready -U postgres

# View database logs
docker-compose logs todo_postgres
```

#### Events not processing
```bash
# Check Kafka consumer logs
docker-compose logs todo_task_service | grep -i kafka

# Verify topic exists
docker exec -it todo_kafka kafka-topics --bootstrap-server localhost:9092 --describe --topic user-events
```

## 🧪 Testing Scenarios

### Scenario 1: User Lifecycle with Tasks
1. Create user → 3 welcome tasks auto-created
2. Add custom tasks for user  
3. Update/complete some tasks
4. Delete user → all tasks auto-deleted

### Scenario 2: Multiple Users
1. Create multiple users
2. Each gets their own welcome tasks
3. Verify task isolation between users
4. Delete one user → only their tasks are removed

### Scenario 3: Service Recovery
1. Stop task service
2. Create/delete users (events queued)  
3. Start task service → events processed
4. Verify data consistency

## 🏁 Stopping the Application

```bash
# Stop all services
docker-compose down

# Stop and remove volumes (⚠️ deletes all data)
docker-compose down -v

# Stop and remove images
docker-compose down --rmi all
```

## 📁 Project Structure

```
todo-microservices/
├── docker-compose.yml           # Multi-service orchestration
├── user-service/               # Django REST API
│   ├── Dockerfile
│   ├── requirements.txt
│   ├── manage.py
│   └── users/
│       ├── views.py            # Updated with Kafka events
│       ├── models.py
│       ├── serializers.py
│       └── kafka_producer.py   # Kafka event producer
├── task-service/               # Go Gin API  
│   ├── Dockerfile
│   ├── go.mod                 # Updated with Kafka dependency
│   ├── main.go                # Updated with Kafka consumer
│   └── kafka_consumer.go      # Kafka event consumer
├── todo-frontend/             # React application
│   ├── Dockerfile
│   └── ... (frontend files)
└── README.md                  # This file
```

## 🎉 Success Indicators

When everything is working correctly, you should see:

1. **All services healthy**: `docker-compose ps` shows all services as "Up"
2. **Auto-task creation**: New users automatically get 3 welcome tasks
3. **Auto-task cleanup**: Deleting users removes all their tasks  
4. **Frontend functional**: React app can create/manage users and tasks
5. **Event processing**: Kafka logs show events being published and consumed
6. **Data consistency**: Database reflects all operations correctly

---

**🎯 This application demonstrates modern microservices patterns including event-driven architecture, containerization, and real-time data synchronization!**