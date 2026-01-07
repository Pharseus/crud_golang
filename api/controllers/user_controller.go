package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Pharseus/crud_golang.git/api/helper"
	"github.com/Pharseus/crud_golang.git/api/payloads"
	"github.com/Pharseus/crud_golang.git/api/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req payloads.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := c.userService.Create(r.Context(), req)
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	helper.RespondSuccess(w, http.StatusCreated, "User created successfully", user)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	req := payloads.PaginationRequest{Page: page, Limit: limit}
	result, err := c.userService.FindAll(r.Context(), req)
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to get users", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "Users retrieved successfully", result)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := c.userService.FindById(r.Context(), int32(id))
	if err != nil {
		helper.RespondError(w, http.StatusNotFound, "User not found", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "User retrieved successfully", user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	var req payloads.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	user, err := c.userService.Update(r.Context(), int32(id), req)
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "User updated successfully", user)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	if err := c.userService.Delete(r.Context(), int32(id)); err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "User deleted successfully", nil)
}
