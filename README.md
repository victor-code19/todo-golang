# 📝 Todo App - Go REST API

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

*A basic Todo application built with Go, Gin framework, and MongoDB featuring both REST API and web interface*

## ✨ Features

### 🔧 Backend Features
- **RESTful API** - Clean and organized REST endpoints
- **Official MongoDB Driver Integration** - Persistent data storage with efficient queries
- **Gin Framework** - Fast HTTP web framework with middleware support
- **JSON API** - Standard JSON responses for all endpoints
- **Error Handling** - Error handling and validation
- **Context Management** - Proper timeout handling for database operations

### 🎨 Frontend Features
- **Interactive UI** - Dynamic web interface with real-time updates
- **Responsive Design** - Mobile-friendly responsive layout
- **AJAX Operations** - Smooth user experience without page reloads
- **Task Management** - Add, delete, and view tasks seamlessly
- **Visual Feedback** - Intuitive user interface with Font Awesome icons

## 🏗️ Project Structure

```
todo-golang/
├── 📁 controllers/         # Business logic and request handlers
│   ├── task.go             # Task controller with CRUD operations
│   └── task_test.go        # Controller unit tests
├── 📁 models/              # Data models and structures
│   ├── task.go             # Task and ViewTask model definitions
│   └── task_test.go        # Model unit tests
├── 📁 public/              # Static assets
│   ├── 📁 css/
│   │   └── style.css       # Application styles
│   ├── 📁 img/
│   │   └── *.ico           # Favicon and images
│   └── 📁 js/
│       └── index.js        # Frontend JavaScript logic
├── 📁 templates/           # HTML templates
│   └── index.gohtml        # Main application template
├── 📄 main.go              # Application entry point and server setup
├── 📄 go.mod               # Go module dependencies
├── 📄 go.sum               # Go module checksums
├── 📄 integration_test.go  # Integration tests
├── 📄 Dockerfile           # Docker container configuration
├── 📄 docker-compose.yml   # Docker Compose setup
├── 📄 Makefile             # Build automation
├── 📄 .air.toml            # Live reload configuration
├── 📄 .env                 # Environment variables
└── 📄 init-mongo.js        # MongoDB initialization script
```

## 🚀 Quick Start

### Prerequisites

Make sure you have the following installed on your system:

- **Docker** (includes Docker Compose) - [Install Docker](https://docs.docker.com/get-docker/)
- **Git** - [Install Git](https://git-scm.com/downloads)

### Installation

1. **Clone the repository**
   ```bash
   git clone <your-repository-url>
   cd todo-golang
   ```

2. **Start the application with Docker Compose**
   ```bash
   docker-compose up --build
   ```

3. **Access the application**
   - **Web Interface**: http://localhost:8080/view/tasks
   - **API Base URL**: http://localhost:8080/api

### Docker Commands

```bash
# Build and start all services
docker-compose up --build

# Run in background (detached mode)
docker-compose up -d

# Stop all services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild and restart
docker-compose down && docker-compose up --build
```

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api
```

### Endpoints

| Method | Endpoint | Description | Request Body | Response |
|--------|----------|-------------|--------------|----------|
| `GET` | `/tasks` | Retrieve all tasks | - | Array of tasks |
| `POST` | `/task` | Create a new task | `{"description": "string"}` | Created task object |
| `DELETE` | `/task/:id` | Delete specific task | - | Success message |
| `DELETE` | `/tasks` | Delete all tasks | - | Success message with count |

### Web Interface

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/view/tasks` | Display tasks in HTML template |

### Example API Usage

#### Create a Task
```bash
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -d '{"description": "Learn Go programming"}'
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "description": "Learn Go programming"
}
```

#### Get All Tasks
```bash
curl http://localhost:8080/api/tasks
```

**Response:**
```json
[
  {
    "id": "507f1f77bcf86cd799439011",
    "description": "Learn Go programming"
  },
  {
    "id": "507f1f77bcf86cd799439012", 
    "description": "Build a REST API"
  }
]
```

#### Delete a Task
```bash
curl -X DELETE http://localhost:8080/api/task/507f1f77bcf86cd799439011
```

**Response:**
```json
{
  "message": "Task deleted successfully"
}
```

#### Delete All Tasks
```bash
curl -X DELETE http://localhost:8080/api/tasks
```

**Response:**
```json
{
  "message": "All tasks deleted successfully",
  "deletedCount": 5
}
```

## 🧪 Testing

### Run Unit Tests
```bash
go test ./...
```

### Run Integration Tests
```bash
go test -tags=integration
```

### Test Coverage
```bash
go test -cover ./...
```

### Using the Web Interface
1. Navigate to http://localhost:8080/view/tasks
2. Add new tasks using the input field
3. Delete individual tasks using the trash icon
4. Clear all tasks using the "Clear all" button

## 🛠️ Technology Stack

### Backend
- **[Go](https://golang.org/)** (v1.25) - Modern programming language
- **[Gin](https://gin-gonic.com/)** - HTTP web framework
- **[MongoDB](https://www.mongodb.com/)** - NoSQL database
- **[MongoDB Go Driver](https://go.mongodb.org/mongo-driver/)** (v2.3.0) - Official MongoDB driver

### Frontend
- **HTML5** - Semantic markup with Go templates
- **CSS3** - Modern styling with Flexbox
- **Vanilla JavaScript** - Interactive functionality with Fetch API
- **[Font Awesome](https://fontawesome.com/)** - Icon library

### Development Tools
- **[Air](https://github.com/cosmtrek/air)** - Live reload for Go apps
- **Docker** - Containerization
- **Make** - Build automation

## 🔧 Configuration

### Environment Variables
- `MONGODB_URI` - MongoDB connection string (default: detected from environment)

### Database Configuration
- **Database**: `todo-app-go`
- **Collection**: `tasks` 
- **Connection Timeout**: 5 seconds

### Database Schema
```json
{
  "_id": "ObjectId",
  "description": "string"
}
```