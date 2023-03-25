package repository

import (
	"context"
	"encoding/json"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/utils"
	"gorm.io/gorm"
)

type UserHistoryRepository struct {
}

type UserHistoryRepositoryInterface interface {
	GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...uint) (interface{}, error)
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

func (repository *UserHistoryRepository) GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...uint) (interface{}, error) {
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
	pagination.SetTotalPages()

	// TODO : set data to json
	users_json, _ := setDataToJson(&users)

	return map[string]interface{}{
		"data":       users_json,
		"pagination": pagination,
	}, nil
}

func setDataToJson(users *[]mysql.UserHistory) (users_json []mysql.UserHistoryRespond, err error) {
	var data map[string]map[string]interface{}
	for i := range *users {
		if err := json.Unmarshal([]byte((*users)[i].Data), &data); err != nil {
			return nil, err
		}
		users_json = append(users_json, mysql.UserHistoryRespond{
			ID:     (*users)[i].ID,
			UserID: (*users)[i].UserID,
			Action: (*users)[i].Action,
			Data: &mysql.UserHistoryData{
				Before: data["before"],
				After:  data["after"],
			},
		})
	}
	return users_json, nil
}
