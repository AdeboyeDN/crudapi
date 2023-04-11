package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Tasks []Task

// Create a slice to store the tasks
var tasks Tasks

func main() {
	// Create a new router with default middleware
	router := gin.Default()

	// Set up routes
	router.GET("/tasks", getTasks)          // Get all tasks
	router.GET("/tasks/:id", getTask)       // Get a specific task by ID
	router.POST("/tasks", createTask)       // Create a new task
	router.PUT("/tasks/:id", updateTask)    // Update a task by ID
	router.DELETE("/tasks/:id", deleteTask) // Delete a task by ID

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

// Get all tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// Get a specific task by ID
func getTask(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Create a new task
func createTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, task)
	c.JSON(http.StatusCreated, task)
}

// Update a task by ID
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = task.ID
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Delete a task by ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
