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

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with name, email, and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body payloads.CreateUserRequest true "Create User Request"
// @Success      201 {object} payloads.SuccessResponse{data=payloads.UserResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/users [post]
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

// GetAll godoc
// @Summary      Get all users
// @Description  Get list of all users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        limit query int false "Items per page" default(10)
// @Success      200 {object} payloads.SuccessResponse{data=payloads.PaginationResponse}
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/users [get]
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

// GetById godoc
// @Summary      Get user by ID
// @Description  Get a single user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} payloads.SuccessResponse{data=payloads.UserResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      404 {object} payloads.ErrorResponse
// @Router       /v1/users/{id} [get]
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

// Update godoc
// @Summary      Update user
// @Description  Update user information by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        request body payloads.UpdateUserRequest true "Update User Request"
// @Success      200 {object} payloads.SuccessResponse{data=payloads.UserResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/users/{id} [put]
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

// Delete godoc
// @Summary      Delete user
// @Description  Soft delete user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} payloads.SuccessResponse
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/users/{id} [delete]
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
