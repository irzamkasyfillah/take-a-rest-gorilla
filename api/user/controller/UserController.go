package controller

import (
	"context"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
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

var validate = validator.New()

type UserControllerInterface interface {
	UserGetHistoryService(ctx context.Context, input *request.UserGetHistoryRequest, write *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format)
	UserGetAllService(ctx context.Context, input *request.UserGetAllRequest, w *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format)
	UserGetByIdService(ctx context.Context, input *request.UserGetOneRequest, w *http.ResponseWriter) *transformer.Format
	UserCreateService(ctx context.Context, input *request.UserCreateRequest, w *http.ResponseWriter) *transformer.Format
	UserUpdateService(ctx context.Context, input *request.UserUpdateRequest, w *http.ResponseWriter) *transformer.Format
	UserDeleteService(ctx context.Context, input *request.UserDeleteRequest, w *http.ResponseWriter) *transformer.Format
}

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		UserService: service,
	}
}

func (controller *UserController) UserGetHistoryService(ctx context.Context, input *request.UserGetHistoryRequest, write *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format) {
	w := *write

	// validate input request
	if err := validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return nil, exception.UserException(
			utils.RequestValidation,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	users, err := controller.UserService.GetHistoryService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
	}
	data := UserTransformer.UserGetAllTransformer(users.(map[string]interface{}))
	return data, nil
}

func (controller *UserController) UserGetAllService(ctx context.Context, input *request.UserGetAllRequest, write *http.ResponseWriter) (*UserTransformer.Format, *transformer.Format) {

	w := *write

	// validate input request
	if err := validate.Struct(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		data := exception.UserException(utils.RequestValidation, map[string]interface{}{"message": err.Error()})
		return nil, data
	}

	users, _ := controller.UserService.GetAllService(ctx, input)
	return UserTransformer.UserGetAllTransformer(users.(map[string]interface{})), nil
}

func (controller *UserController) UserGetByIdService(ctx context.Context, input *request.UserGetOneRequest, write *http.ResponseWriter) *transformer.Format {
	w := *write

	// validate input request
	if err := validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(
			utils.RequestValidation,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	user, err := controller.UserService.GetUserService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserCreateService(ctx context.Context, input *request.UserCreateRequest, write *http.ResponseWriter) *transformer.Format {
	w := *write

	// validate input request
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(
			utils.RequestValidation,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	user, err := controller.UserService.CreateService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserUpdateService(ctx context.Context, input *request.UserUpdateRequest, write *http.ResponseWriter) *transformer.Format {
	w := *write

	// validate input request
	if err := validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(
			utils.RequestValidation,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	// convert request struct to map
	input_request := make(map[string]interface{})
	t := reflect.TypeOf(*input)
	v := reflect.ValueOf(*input)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		value := v.Field(i).Interface()
		if value != "" {
			input_request[tag] = value
		}
	}

	user, err := controller.UserService.UpdateService(ctx, input_request)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(user)
}

func (controller *UserController) UserDeleteService(ctx context.Context, input *request.UserDeleteRequest, write *http.ResponseWriter) *transformer.Format {
	w := *write

	// validate input request
	if err := validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return exception.UserException(
			utils.RequestValidation,
			map[string]interface{}{
				"message": err.Error(),
			},
		)
	}

	err := controller.UserService.DeleteService(ctx, input)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		return exception.UserException(err.Message, err.Data)
	}
	return UserTransformer.UserTransformer(map[string]interface{}{"message": "Data has been deleted"})
}
