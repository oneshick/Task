package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskService.NewTRepository(database)
	tService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(tService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.CreateTask)
	e.PATCH("/tasks/:id", taskHandlers.UpdateTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
