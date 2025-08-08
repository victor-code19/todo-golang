
# 📝 Todo App - Go REST API

![Go](https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

*Basic Todo application built with Go, Gin framework, and MongoDB*

[🚀 Demo](https://youtu.be/8Tl3OFXLrAo) • [📖 Documentation](#api-documentation) • [🛠️ Installation](#installation)

## ✨ Features

### 🔧 Backend Features
- **RESTful API** - Clean and organized REST endpoints
- **MongoDB Integration** - Persistent data storage with efficient queries
- **Gin Framework** - Fast HTTP web framework with middleware support
- **JSON API** - Standard JSON responses for all endpoints
- **Error Handling** - Comprehensive error handling and validation

### 🎨 Frontend Features
- **Interactive UI** - Dynamic web interface with real-time updates
- **Responsive Design** - Mobile-friendly responsive layout
- **AJAX Operations** - Smooth user experience without page reloads
- **Task Management** - Add, delete, and view tasks seamlessly
- **Visual Feedback** - Intuitive user interface with Font Awesome icons

## 🏗️ Project Structure

```
todo-golang/
├── 📁 controllers/          # Business logic and request handlers
│   └── task.go             # Task controller with CRUD operations
├── 📁 models/              # Data models and structures
│   └── task.go             # Task model definition
├── 📁 public/              # Static assets
│   ├── 📁 css/
│   │   └── style.css       # Application styles
│   ├── 📁 img/
│   │   └── *.ico           # Favicon and images
│   └── 📁 js/
│       └── index.js        # Frontend JavaScript logic
├── 📁 templates/           # HTML templates
│   └── index.gohtml        # Main application template
├── 📄 main.go              # Application entry point
├── 📄 go.mod               # Go module dependencies
├── 📄 go.sum               # Go module checksums
└── 📄 README.md            # This file
```

## 🚀 Quick Start

### Prerequisites

Make sure you have the following installed on your system:

- **Go** (version 1.17 or higher) - [Download Go](https://golang.org/dl/)
- **MongoDB** (version 4.0 or higher) - [Install MongoDB](https://docs.mongodb.com/manual/installation/)
- **Git** - [Install Git](https://git-scm.com/downloads)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/victor-code19/todo-golang.git
   cd todo-golang
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Start MongoDB service**
   ```bash
   # On Linux/macOS
   sudo systemctl start mongod
   
   # On macOS with Homebrew
   brew services start mongodb-community
   
   # On Windows
   net start MongoDB
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

5. **Access the application**
   - **Web Interface**: http://localhost:8080/view/tasks
   - **API Base URL**: http://localhost:8080/api

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
| `DELETE` | `/tasks` | Delete all tasks | - | Success message |

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

#### Get All Tasks
```bash
curl http://localhost:8080/api/tasks
```

#### Delete a Task
```bash
curl -X DELETE http://localhost:8080/api/task/{task-id}
```

## 🧪 Testing the Application

### Using the Web Interface
1. Navigate to http://localhost:8080/view/tasks
2. Add new tasks using the input field
3. Delete individual tasks using the trash icon
4. Clear all tasks using the "Clear all" button

### Using API with curl
```bash
# Test API endpoints
./test-api.sh  # (create this script for automated testing)
```

## 🛠️ Technology Stack

### Backend
- **[Go](https://golang.org/)** - Modern programming language
- **[Gin](https://gin-gonic.com/)** - HTTP web framework
- **[MongoDB](https://www.mongodb.com/)** - NoSQL database
- **[mgo](https://github.com/go-mgo/mgo)** - MongoDB driver for Go

### Frontend
- **HTML5** - Semantic markup
- **CSS3** - Modern styling with Flexbox
- **Vanilla JavaScript** - Interactive functionality
- **[Font Awesome](https://fontawesome.com/)** - Icon library

## 🔧 Configuration

### Database Configuration
The application connects to MongoDB on the default port:
```go
mongodb://127.0.0.1:27017
```

### Database Schema
- **Database**: `todo-app-go`
- **Collection**: `tasks`
- **Document Structure**:
  ```json
  {
    "_id": "ObjectId",
    "description": "string"
  }
  ```

## 🎥 Demo

Check out the application in action: [YouTube Demo](https://youtu.be/8Tl3OFXLrAo)