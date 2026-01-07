package services

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
	"github.com/Pharseus/crud_golang.git/api/payloads"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	repositoryImpl "github.com/Pharseus/crud_golang.git/api/repositories/RepositoryImpl"
	"github.com/Pharseus/crud_golang.git/api/securities"
	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx context.Context, req payloads.CreateUserRequest) (payloads.UserResponse, error)
	FindAll(ctx context.Context, pagination payloads.PaginationRequest) (payloads.PaginationResponse, error)
	FindById(ctx context.Context, id int32) (payloads.UserResponse, error)
	Update(ctx context.Context, id int32, req payloads.UpdateUserRequest) (payloads.UserResponse, error)
	Delete(ctx context.Context, id int32) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		userRepo: repositoryImpl.NewUserRepositoryImpl(db),
	}
}

func (s *userServiceImpl) Create(ctx context.Context, req payloads.CreateUserRequest) (payloads.UserResponse, error) {
	hashedPassword, err := securities.HashPassword(req.Password)
	if err != nil {
		return payloads.UserResponse{}, err
	}

	user := entities.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		IsActive:     true,
	}

	created, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return payloads.UserResponse{}, err
	}

	return payloads.UserResponse{
		Id:        created.Id,
		Name:      created.Name,
		Email:     created.Email,
		IsActive:  created.IsActive,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (s *userServiceImpl) FindAll(ctx context.Context, req payloads.PaginationRequest) (payloads.PaginationResponse, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit
	pagination := repository.Pagination{Limit: req.Limit, Offset: offset}

	users, total, err := s.userRepo.FindUser(ctx, pagination)
	if err != nil {
		return payloads.PaginationResponse{}, err
	}

	var userResponses []payloads.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, payloads.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	totalPages := int(total) / req.Limit
	if int(total)%req.Limit != 0 {
		totalPages++
	}

	return payloads.PaginationResponse{
		Data:       userResponses,
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}, nil
}

func (s *userServiceImpl) FindById(ctx context.Context, id int32) (payloads.UserResponse, error) {
	user, err := s.userRepo.FindUserById(ctx, id)
	if err != nil {
		return payloads.UserResponse{}, err
	}

	return payloads.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userServiceImpl) Update(ctx context.Context, id int32, req payloads.UpdateUserRequest) (payloads.UserResponse, error) {
	user := entities.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if req.Password != "" {
		hashedPassword, err := securities.HashPassword(req.Password)
		if err != nil {
			return payloads.UserResponse{}, err
		}
		user.PasswordHash = hashedPassword
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	updated, err := s.userRepo.UpdateUserById(ctx, user, id)
	if err != nil {
		return payloads.UserResponse{}, err
	}

	return payloads.UserResponse{
		Id:        updated.Id,
		Name:      updated.Name,
		Email:     updated.Email,
		IsActive:  updated.IsActive,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}, nil
}

func (s *userServiceImpl) Delete(ctx context.Context, id int32) error {
	return s.userRepo.DeleteUserById(ctx, id)
}
