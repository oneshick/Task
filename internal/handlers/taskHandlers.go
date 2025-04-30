package handlers

import (
	"Tasks/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task, err := h.service.CreateTask(req.Title, req.Status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error creating task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedTask, err := h.service.UpdateTask(id, req.Title, req.Status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error updating task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting task"})
	}
	return c.NoContent(http.StatusNoContent)
}
