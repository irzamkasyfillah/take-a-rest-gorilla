package service

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/service/action"
)

type UserService struct {
	UserAction action.UserActionInterface
}

type UserServiceInterface interface {
	GetHistoryService(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError)
	GetAllService(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError)
	GetUserService(ctx context.Context, input request.UserGetOneModel) (*mysql.User, *exception.HandleError)
	CreateService(ctx context.Context, input *mysql.User) (*mysql.User, *exception.HandleError)
	UpdateService(ctx context.Context, input map[string]interface{}) (*mysql.User, *exception.HandleError)
	DeleteService(ctx context.Context, input request.UserGetOneModel) *exception.HandleError
}

func NewUserService(action action.UserActionInterface) UserServiceInterface {
	return &UserService{
		UserAction: action,
	}
}

func (service *UserService) GetHistoryService(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError) {
	return service.UserAction.UserGetHistoryAction(ctx, input)
}

func (service *UserService) GetAllService(ctx context.Context, input map[string]interface{}) (interface{}, *exception.HandleError) {
	return service.UserAction.UserGetAllAction(ctx, input)
}

func (service *UserService) GetUserService(ctx context.Context, input request.UserGetOneModel) (*mysql.User, *exception.HandleError) {
	return service.UserAction.UserGetOneAction(ctx, input)
}

func (service *UserService) CreateService(ctx context.Context, input *mysql.User) (*mysql.User, *exception.HandleError) {
	return service.UserAction.UserCreateAction(ctx, input)
}

func (service *UserService) UpdateService(ctx context.Context, input map[string]interface{}) (*mysql.User, *exception.HandleError) {
	return service.UserAction.UserUpdateAction(ctx, input)
}

func (service *UserService) DeleteService(ctx context.Context, input request.UserGetOneModel) *exception.HandleError {
	return service.UserAction.UserDeleteAction(ctx, input)
}
