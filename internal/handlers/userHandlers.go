package handlers

import (
	"Tasks/internal/userService"
	"Tasks/internal/web/users"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var response users.GetUsers200JSONResponse
	for _, u := range allUsers {
		response = append(response, users.User{
			Id:    &u.ID,
			Email: &u.Email,
		})
	}
	return response, nil
}

func (h *UserHandler) PostUser(_ context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	userRequest := request.Body

	UserToCreate := userService.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	createdUser, err := h.service.CreateUser(UserToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUser201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}

	return response, nil
}

func (h *UserHandler) DeleteUserByID(_ context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteUser(id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *UserHandler) PatchUserByID(_ context.Context, request users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	user, err := h.service.GetUserByID(request.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.PatchUserByID404Response{}, nil
		}
		return nil, err
	}

	if request.Body.Email != nil {
		user.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		user.Password = *request.Body.Password
	}

	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	response := users.User{
		Id:    &updatedUser.ID,
		Email: &updatedUser.Email,
	}
	return users.PatchUserByID200JSONResponse(response), nil
}
