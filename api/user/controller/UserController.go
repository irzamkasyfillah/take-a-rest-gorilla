package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/service"
	UserTransformer "github.com/irzam/my-app/api/user/transformer"
	"github.com/irzam/my-app/api/user/utils"
	"github.com/irzam/my-app/lib/transformer"
)

type UserController struct {
	UserService service.UserServiceInterface
}

type UserControllerInterface interface {
	UserGetHistoryService(ctx context.Context, input map[string]interface{}, write *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format)
	UserGetAllService(ctx context.Context, input map[string]interface{}, w *http.ResponseWriter) UserTransformer.Format
	UserGetByIdService(ctx context.Context, input map[string]string, w *http.ResponseWriter) transformer.Format
	UserCreateService(ctx context.Context, input *mysql.User, w *http.ResponseWriter) transformer.Format
	UserUpdateService(ctx context.Context, input map[string]interface{}, w *http.ResponseWriter) transformer.Format
	UserDeleteService(ctx context.Context, input map[string]string, w *http.ResponseWriter) transformer.Format
}

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		UserService: service,
	}
}

func (controller *UserController) UserGetHistoryService(ctx context.Context, input map[string]interface{}, write *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format) {
	w := *write
	// Check if id is valid (is a number)
	user_id, _ := input["user_id"].(float64)
	if user_id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := exception.UserException(utils.UserNotFound, map[string]interface{}{"id": input["user_id"]})
		return nil, &data
	}
	users, err := controller.UserService.GetHistoryService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
	}
	data := UserTransformer.UserGetAllTransformer(users.(map[string]interface{}))
	return &data, nil
}

func (controller *UserController) UserGetAllService(ctx context.Context, input map[string]interface{}, write *http.ResponseWriter) UserTransformer.Format {
	users, _ := controller.UserService.GetAllService(ctx, input)
	return UserTransformer.UserGetAllTransformer(users.(map[string]interface{}))
}

func (controller *UserController) UserGetByIdService(ctx context.Context, input map[string]string, write *http.ResponseWriter) transformer.Format {
	w := *write
	// Check if id is valid (is a number)
	id, _ := strconv.Atoi(input["id"])
	if id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(utils.UserNotFound, map[string]interface{}{"id": input["id"]})
	}

	user, err := controller.UserService.GetUserService(ctx, request.UserGetOneModel{ID: uint(id)})
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserCreateService(ctx context.Context, input *mysql.User, write *http.ResponseWriter) transformer.Format {
	w := *write
	user, err := controller.UserService.CreateService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserUpdateService(ctx context.Context, input map[string]interface{}, write *http.ResponseWriter) transformer.Format {
	w := *write
	// Check if id is valid (is a number)
	id, _ := strconv.Atoi(input["id"].(string))
	if id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(utils.UserNotFound, map[string]interface{}{"id": input["id"]})
	}

	user, err := controller.UserService.UpdateService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserDeleteService(ctx context.Context, input map[string]string, write *http.ResponseWriter) transformer.Format {
	w := *write
	// Check if id is valid (is a number)
	id, _ := strconv.Atoi(input["id"])
	if id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(utils.UserNotFound, map[string]interface{}{"id": input["id"]})
	}

	err := controller.UserService.DeleteService(ctx, request.UserGetOneModel{ID: uint(id)})
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(map[string]interface{}{"message": "Data has been deleted"})
}
