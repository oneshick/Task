package handlers

import (
	"Tasks/internal/taskService"
	"Tasks/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, t := range allTasks {
		task := tasks.Task{
			Id:     &t.ID,
			Title:  &t.Title,
			Status: &t.Status,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Title:  taskRequest.Title,
		Status: taskRequest.Status,
	}

	createdTask, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Title:  &createdTask.Title,
		Status: &createdTask.Status,
	}

	return response, nil
}

func (h *TaskHandler) PatchTask(_ context.Context, req tasks.PatchTaskRequestObject) (tasks.PatchTaskResponseObject, error) {
	id := req.Id
	body := req.Body

	task, err := h.service.GetTaskById(id)
	if err != nil {
		return tasks.PatchTask404Response{}, nil
	}

	if body.Title != nil {
		task.Title = *body.Title
	}
	if body.Status != nil {
		task.Status = *body.Status
	}

	updatedTask, err := h.service.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	return tasks.PatchTask200JSONResponse{
		Id:     &updatedTask.ID,
		Title:  &updatedTask.Title,
		Status: &updatedTask.Status,
	}, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteTask(id); err != nil {
		return nil, err
	}
	return nil, nil
}
