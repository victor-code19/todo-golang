
## Todo REST API with Gin Framework and MongoDB

This project demonstrates a simple Todo REST API using the Gin web framework and MongoDB for data storage. It allows creating, fetching, and deleting tasks through API endpoints and also provides a basic HTML view to display tasks. Provides a clean and organized structure for API controllers.

## Table of Contents

* [Features](#features)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
* [Endpoints](#endpoints)
* [Demo](#demo)

## Features
* Fetch tasks from MongoDB and respond with JSON data.
* Create new tasks and store them in the database.
* Delete tasks by ID and delete all tasks.
* Display tasks in an HTML template.
* Utilizes the Gin web framework for routing and handling HTTP requests.

## Prerequisites
* Go programming language (at least Go 1.13).
* MongoDB server running locally.


## Installation
Clone the repository: 
        
    git clone https://github.com/victor-code19/todo-golang.git

Navigate to the project directory and install the required dependencies:
    
    go get github.com/gin-gonic/gin
    go get gopkg.in/mgo.v2

## Usage

Ensure your MongoDB server is up and running.

Run the application:

    go run main.go

Access the web interface by navigating to http://localhost:8080/ in your browser.

## Endpoints

* GET /tasks - Fetch all tasks as JSON.
* POST /tasks - Create a new task.
* DELETE /tasks/:id - Delete a task by ID.
* GET /tasks/all - Display all tasks in an HTML template.
* DELETE /tasks/all - Delete all tasks.

## Demo 
    https://youtu.be/8Tl3OFXLrAo


