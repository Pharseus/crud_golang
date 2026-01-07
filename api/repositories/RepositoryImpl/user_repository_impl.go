package repositoryImpl

import (
	"context"
	"fmt"

	"github.com/Pharseus/crud_golang.git/api/entities"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: db}
}
func (repository *userRepositoryImpl) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	result := repository.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
	// script := "INSERT INTO users (name,email,passwordhash,isactive) VALUES (?,?,?,?)"
	// result, err := repository.DB.Create(ctx, script, user.Name, user.Email, user.PasswordHash, user.IsActive)
	// if err != nil {
	// 	return user, nil
	// }
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return user, nil
	// }
	// user.Id = int(id)
	// return user, nil
}
func (repository *userRepositoryImpl) FindUser(ctx context.Context, pagination repository.Pagination) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	if err := repository.DB.WithContext(ctx).Model(&entities.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := repository.DB.WithContext(ctx).Limit(pagination.Limit).Offset(pagination.Offset).Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}

func (repository *userRepositoryImpl) FindUserById(ctx context.Context, id int32) (entities.User, error) {
	var users entities.User

	result := repository.DB.WithContext(ctx).Where("id = ?", id).First(&users)

	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

// UpdateUserById(ctx context.Context, id int32) (entities.User, error)
func (repository *userRepositoryImpl) UpdateUserById(ctx context.Context, user entities.User, id int32) (entities.User, error) {
	result := repository.DB.WithContext(ctx).Model(&user).Where("id = ?", id).Updates(user)

	if result.Error != nil {
		return entities.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return entities.User{}, gorm.ErrRecordNotFound
	}

	var updatedUser entities.User
	err := repository.DB.WithContext(ctx).First(&updatedUser, id).Error
	if err != nil {
		return entities.User{}, err
	}
	return updatedUser, nil
}

// DeleteUserById(ctx context.Context, id int32) error
func (repository *userRepositoryImpl) DeleteUserById(ctx context.Context, id int32) error {
	var user entities.User
	result := repository.DB.WithContext(ctx).Model(&user).Where("id = ?", id).Delete(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	fmt.Printf("Data Berhasil di hapus, id %d", id)
	return nil
}
