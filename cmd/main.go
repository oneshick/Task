package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"Tasks/internal/userService"
	"Tasks/internal/web/tasks"
	"Tasks/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Tasks setup
	tasksRepo := taskService.NewTRepository(database)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandlers := handlers.NewTaskHandler(tasksService)

	// Users setup
	usersRepo := userService.NewUserRepository(database)
	usersService := userService.NewUserService(usersRepo)
	usersHandlers := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Register task routes
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandlers, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Register user routes
	usersStrictHandler := users.NewStrictHandler(usersHandlers, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
