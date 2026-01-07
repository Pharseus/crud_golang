package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/Pharseus/crud_golang.git/api/config"
	"github.com/Pharseus/crud_golang.git/api/entities"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	repositoryImpl "github.com/Pharseus/crud_golang.git/api/repositories/RepositoryImpl"
	"github.com/Pharseus/crud_golang.git/api/securities"
)

func TestCreateUser(t *testing.T) {
	pass, err := securities.HashPassword("Hayolo")
	if err != nil {
		panic(err)
	}
	cfg := config.LoadConfig()
	userRepository := repositoryImpl.NewUserRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	user := entities.User{
		Name:         "NewUser",
		Email:        "NewUser@gmail.com",
		PasswordHash: pass,
		IsActive:     true,
	}
	result, err := userRepository.CreateUser(ctx, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindUser(t *testing.T) {
	var users []entities.User
	var total int64
	cfg := config.LoadConfig()
	userRepository := repositoryImpl.NewUserRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	usrPagination := repository.Pagination{Limit: 2, Offset: 0}
	users, total, err := userRepository.FindUser(ctx, usrPagination)
	if err != nil {
		panic(err)
	}
	fmt.Println(users, total)
}

func TestFindUserById(t *testing.T) {
	var users entities.User
	var id int32
	id = 3

	cfg := config.LoadConfig()
	userRepository := repositoryImpl.NewUserRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	users, err := userRepository.FindUserById(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}

func TestUpdateUser(t *testing.T) {
	pass, err := securities.HashPassword("ridwan123")
	if err != nil {
		panic(err)
	}
	user := entities.User{
		PasswordHash: pass,
	}
	var id int32
	id = 2

	cfg := config.LoadConfig()
	userRepository := repositoryImpl.NewUserRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	user, err = userRepository.UpdateUserById(ctx, user, id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Update data di id-%d, Berhasil", id)
	fmt.Println(user)
}

func TestDeleteUser(t *testing.T) {
	id := 7

	cfg := config.LoadConfig()
	userRepository := repositoryImpl.NewUserRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	err := userRepository.DeleteUserById(ctx, int32(id))
	if err != nil {
		panic(err)
	}
}
