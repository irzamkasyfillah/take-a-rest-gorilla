package action

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/entity/repository"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"gorm.io/gorm"
)

type UserAction struct {
	UserHistoryRepository repository.UserHistoryRepositoryInterface
	UserRepository        repository.UserRepositoryInterface
	DB                    *gorm.DB
}

type UserActionInterface interface {
	UserGetAllAction(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError)
	UserGetOneAction(ctx context.Context, input request.UserGetOneModel) (*mysql.User, *exception.HandleError)
	UserCreateAction(ctx context.Context, input *mysql.User) (*mysql.User, *exception.HandleError)
	UserUpdateAction(ctx context.Context, input map[string]interface{}) (*mysql.User, *exception.HandleError)
	UserDeleteAction(ctx context.Context, input request.UserGetOneModel) *exception.HandleError
	UserGetHistoryAction(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError)
}

func NewUserAction(userRepo repository.UserRepositoryInterface, userHistoryRepo repository.UserHistoryRepositoryInterface, db *gorm.DB) UserActionInterface {
	return &UserAction{
		UserRepository:        userRepo,
		UserHistoryRepository: userHistoryRepo,
		DB:                    db,
	}
}
