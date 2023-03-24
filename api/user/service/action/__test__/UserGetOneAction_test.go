package __test__

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/utils"
	"github.com/stretchr/testify/assert"
)

// unit test + without connection to database
func TestUserGetOneActionMock(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *request.UserGetOneRequest
		Want    *mysql.User
		WantErr *exception.HandleError
	}{
		{
			Name: "Test User Get By ID Action",
			Input: &request.UserGetOneRequest{
				ID: 1,
			},
			Want: &mysql.User{
				ID:       1,
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			WantErr: nil,
		},
		{
			Name: "Test User Get By ID Action (User not found)",
			Input: &request.UserGetOneRequest{
				ID: 999,
			},
			Want: nil,
			WantErr: &exception.HandleError{
				Message:    utils.UserNotFound,
				StatusCode: http.StatusUnprocessableEntity,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			UserRepository.Mock.On("GetByID", context.Background(), tt.Input.ID).Return(tt.Want, nil)
			result, err := UserAction.UserGetOneAction(context.Background(), tt.Input)
			if result != nil {
				assert.Equal(t, tt.Want.Name, result.Name, fmt.Sprintf("got %v, want %v", result.Name, tt.Want.Name))
				assert.Equal(t, tt.Want.Email, result.Email, fmt.Sprintf("got %v, want %v", result.Email, tt.Want.Email))
				assert.Equal(t, tt.Want.Password, result.Password, fmt.Sprintf("got %v, want %v", result.Password, tt.Want.Password))
			}
			if err != nil {
				assert.Equal(t, err.Message, tt.WantErr.Message, fmt.Sprintf("got %v, want %v", err.Message, tt.WantErr.Message))
				assert.Equal(t, err.StatusCode, tt.WantErr.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, tt.WantErr.StatusCode))
			}
		})
		UserRepository.Mock.ExpectedCalls = nil
	}
}

// unit test + connection to database
// func TestUserGetOneAction(t *testing.T) {
// 	userHistoryRepository := repository.NewUserHistoryRepository()
// 	userRepository := repository.NewUserRepository()
// 	userAction := action.NewUserAction(userRepository, userHistoryRepository, SetupDB())

// 	ctx := context.Background()

// 	result, _ := userAction.UserCreateAction(ctx, &mysql.User{
// 		Name:     "Test",
// 		Email:    "tes@gmail.com",
// 		Password: "123456",
// 	})

// 	tests := []struct {
// 		Name    string
// 		ID      uint
// 		Want    *mysql.User
// 		WantErr *exception.HandleError
// 	}{
// 		{
// 			Name: "Test User 12 Create Action",
// 			ID:   result.ID,
// 			Want: &mysql.User{
// 				ID:       result.ID,
// 				Name:     "Test",
// 				Email:    "tes@gmail.com",
// 				Password: "123456",
// 			},
// 			WantErr: nil,
// 		},
// 		{
// 			Name: "Test User 13 Create Action (User not found)",
// 			ID:   9999,
// 			Want: nil,
// 			WantErr: &exception.HandleError{
// 				Message:    utils.UserNotFound,
// 				Data:       map[string]interface{}{"id": 9999},
// 				StatusCode: http.StatusUnprocessableEntity,
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {
// 			result, err := userAction.UserGetOneAction(ctx, request.UserGetOneModel{ID: tt.ID})
// 			if result != nil {
// 				assert.Equal(t, tt.Want.Name, result.Name, fmt.Sprintf("got %v, want %v", result.Name, tt.Want.Name))
// 				assert.Equal(t, tt.Want.Email, result.Email, fmt.Sprintf("got %v, want %v", result.Email, tt.Want.Email))
// 				assert.Equal(t, tt.Want.Password, result.Password, fmt.Sprintf("got %v, want %v", result.Password, tt.Want.Password))
// 			}
// 			if err != nil {
// 				assert.Equal(t, err.Message, tt.WantErr.Message, fmt.Sprintf("got %v, want %v", err.Message, tt.WantErr.Message))
// 				assert.Equal(t, err.StatusCode, tt.WantErr.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, tt.WantErr.StatusCode))
// 			}
// 		})
// 	}

// 	err := userAction.UserDeleteAction(ctx, request.UserGetOneModel{ID: result.ID})
// 	if err != nil {
// 		t.Errorf("UserDeleteAction() error = %v", err)
// 	}
// }
