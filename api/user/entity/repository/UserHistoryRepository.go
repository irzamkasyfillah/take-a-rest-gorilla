package repository

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/utils"
	"gorm.io/gorm"
)

type UserHistoryRepository struct {
}

type UserHistoryRepositoryInterface interface {
	GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...int) (interface{}, error)
	Create(ctx context.Context, db *gorm.DB, userHistory *mysql.UserHistory) (*mysql.UserHistory, error)
}

func NewUserHistoryRepository() UserHistoryRepositoryInterface {
	return &UserHistoryRepository{}
}

func (repository *UserHistoryRepository) Create(ctx context.Context, db *gorm.DB, userHistory *mysql.UserHistory) (*mysql.UserHistory, error) {
	db = db.WithContext(ctx)

	if err := db.Table(mysql.UserHistoryTableName()).Create(userHistory).Error; err != nil {
		return nil, err
	}
	return userHistory, nil
}

func (repository *UserHistoryRepository) GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...int) (interface{}, error) {
	db = db.WithContext(ctx)

	// Default pagination
	var currentPage, perPage int
	if len(args) > 0 {
		currentPage = args[0]
		perPage = args[1]
	} else {
		currentPage = 1
		perPage = 10
	}

	var users []mysql.UserHistory
	pagination := utils.Pagination{
		PerPage:     perPage,
		CurrentPage: currentPage,
	}

	if err := db.Table(mysql.UserHistoryTableName()).
		Where("user_id = ?", user_id).
		Count(&pagination.Total).
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&users).Error; err != nil {
		return nil, err
	}
	// TODO : set data to json
	// for i := range users {
	// 	js, _ := json.Marshal(users[i].Data)
	// 	users[i].Data = js
	// }
	pagination.SetTotalPages()
	return map[string]interface{}{
		"data":       users,
		"pagination": pagination,
	}, nil
}
