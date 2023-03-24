package repository

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
}

type UserRepositoryInterface interface {
	GetByEmail(ctx context.Context, db *gorm.DB, email string) (user *mysql.User, err error)
	GetByID(ctx context.Context, db *gorm.DB, id uint) (user *mysql.User, err error)
	GetAll(ctx context.Context, db *gorm.DB, args ...uint) (interface{}, error)
	Create(ctx context.Context, db *gorm.DB, input *request.UserCreateRequest) (user *mysql.User, err error)
	Update(ctx context.Context, db *gorm.DB, input map[string]interface{}) (user *mysql.User, err error)
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) GetByEmail(ctx context.Context, db *gorm.DB, email string) (user *mysql.User, err error) {
	db = db.WithContext(ctx)

	if err := db.Table(mysql.UserTableName()).Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) GetByID(ctx context.Context, db *gorm.DB, id uint) (user *mysql.User, err error) {
	db = db.WithContext(ctx)

	if err := db.Table(mysql.UserTableName()).Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) GetAll(ctx context.Context, db *gorm.DB, args ...uint) (interface{}, error) {
	db = db.WithContext(ctx)

	// Default pagination
	var currentPage, perPage int
	if len(args) > 0 {
		currentPage = int(args[0])
		perPage = int(args[1])
	} else {
		currentPage = 1
		perPage = 10
	}
	var users []mysql.User
	pagination := utils.Pagination{
		PerPage:     perPage,
		CurrentPage: currentPage,
	}
	if err := db.Table(mysql.UserTableName()).
		Count(&pagination.Total).
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&users).Error; err != nil {
		return nil, err
	}
	pagination.SetTotalPages()
	return map[string]interface{}{
		"data":       users,
		"pagination": pagination,
	}, nil
}

func (repository *UserRepository) Create(ctx context.Context, db *gorm.DB, input *request.UserCreateRequest) (user *mysql.User, err error) {
	// generate input
	user = &mysql.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	db = db.WithContext(ctx)
	if err := db.Table(mysql.UserTableName()).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, db *gorm.DB, input map[string]interface{}) (user *mysql.User, err error) {
	db = db.WithContext(ctx)

	// Update
	if err := db.Table(mysql.UserTableName()).Where("id = ?", input["id"]).Updates(input).Error; err != nil {
		return nil, err
	}
	// Get updated data
	if err := db.Table(mysql.UserTableName()).Where("id = ?", input["id"]).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) Delete(ctx context.Context, db *gorm.DB, id uint) error {
	db = db.WithContext(ctx)

	if err := db.Table(mysql.UserTableName()).Where("id = ?", id).Delete(&mysql.User{}).Error; err != nil {
		return err
	}
	return nil
}
