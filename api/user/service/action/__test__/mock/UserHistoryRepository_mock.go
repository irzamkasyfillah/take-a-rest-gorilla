package mock

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserHistoryRepositoryMock struct {
	Mock mock.Mock
}

type UserHistoryRepositoryInterface interface {
	GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...int) (interface{}, error)
	Create(ctx context.Context, db *gorm.DB, userHistory *mysql.UserHistory) (*mysql.UserHistory, error)
}

func (u *UserHistoryRepositoryMock) GetByUserID(ctx context.Context, db *gorm.DB, user_id uint, args ...int) (interface{}, error) {
	var currentPage, perPage int
	if len(args) > 0 {
		currentPage = args[0]
		perPage = args[1]
	} else {
		currentPage = 1
		perPage = 10
	}
	arg := u.Mock.Called(ctx, user_id, currentPage, perPage)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		return arg.Get(0), arg.Error(1)
	}
}

func (u *UserHistoryRepositoryMock) Create(ctx context.Context, db *gorm.DB, userHistory *mysql.UserHistory) (*mysql.UserHistory, error) {
	arg := u.Mock.Called(ctx, userHistory)
	if arg.Get(0) == nil {
		return nil, arg.Error(0)
	} else {
		return arg.Get(0).(*mysql.UserHistory), arg.Error(1)
	}
}
