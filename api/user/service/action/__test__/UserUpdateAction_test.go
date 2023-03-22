package __test__

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// unit test without database
func TestUserUpdateActionMock(t *testing.T) {
	ctx := context.Background()

	// mock data for testing
	tests := []struct {
		name       string
		input      map[string]interface{}
		userExist  *mysql.User
		emailExist *mysql.User
		want       *mysql.User
		wantErr    *exception.HandleError
	}{
		{
			name: "Test User Update Action",
			input: map[string]interface{}{
				"id":       "1",
				"name":     "Test 123",
				"email":    "tes@gmail.com",
				"password": "123456789",
			},
			userExist: &mysql.User{
				ID:       1,
				Name:     "Test",
				Email:    "tes123@gmail.com",
				Password: "123456",
			},
			emailExist: nil,
			want: &mysql.User{
				ID:       1,
				Name:     "Test 123",
				Email:    "tes@gmail.com",
				Password: "123456789",
			},
			wantErr: nil,
		},
		{
			name: "Test User Update Action (Email already exist)",
			input: map[string]interface{}{
				"id":       "1",
				"name":     "Test 123",
				"email":    "tes@gmail.com",
				"password": "123456789",
			},
			userExist: &mysql.User{
				ID:       1,
				Name:     "Test",
				Email:    "tes123@gmail.com",
				Password: "123456",
			},
			emailExist: &mysql.User{
				ID:       2,
				Name:     "Test 2",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			want: nil,
			wantErr: &exception.HandleError{
				Message:    utils.EmailAlreadyExist,
				StatusCode: http.StatusUnprocessableEntity,
			},
		},
		{
			name: "Test User Update Action (User not found)",
			input: map[string]interface{}{
				"id":       "1",
				"name":     "Test 123",
				"email":    "tes@gmail.com",
				"password": "123456789",
			},
			userExist:  nil,
			emailExist: nil,
			want:       nil,
			wantErr: &exception.HandleError{
				Message:    utils.UserNotFound,
				StatusCode: http.StatusUnprocessableEntity,
			},
		},
	}

	// run test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, _ := strconv.Atoi(test.input["id"].(string))
			UserRepository.Mock.On("GetByID", ctx, uint(id)).Return(test.userExist, nil)
			UserRepository.Mock.On("GetByEmail", ctx, test.input["email"].(string)).Return(test.emailExist, nil)
			UserHistoryRepository.Mock.On("Create", ctx, mock.Anything).Return(nil)
			UserRepository.Mock.On("Update", ctx, test.input).Return(test.want, test.wantErr)

			result, err := UserAction.UserUpdateAction(ctx, test.input)
			if result != nil {
				assert.Equal(t, result.Name, test.want.Name, fmt.Sprintf("got %v, want %v", result.Name, test.want.Name))
				assert.Equal(t, result.Email, test.want.Email, fmt.Sprintf("got %v, want %v", result.Email, test.want.Email))
				assert.Equal(t, result.Password, test.want.Password, fmt.Sprintf("got %v, want %v", result.Password, test.want.Password))
			}
			if err != nil {
				assert.Equal(t, err.Message, test.wantErr.Message, fmt.Sprintf("got %v, want %v", err.Message, test.wantErr.Message))
				assert.Equal(t, err.StatusCode, test.wantErr.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, test.wantErr.StatusCode))
			}
		})
		UserRepository.Mock.ExpectedCalls = nil
	}
}

// unit test with connection to database
// func TestUserUpdateAction(t *testing.T) {
// 	userHistoryRepository := repository.NewUserHistoryRepository()
// 	userRepository := repository.NewUserRepository()
// 	userAction := action.NewUserAction(userRepository, userHistoryRepository, SetupDB())

// 	// Create a new context
// 	ctx := context.Background()

// 	// Create new data for testing
// 	result, _ := userAction.UserCreateAction(ctx, &mysql.User{
// 		Name:     "Test",
// 		Email:    "tes@gmail.com",
// 		Password: "123456",
// 	})

// 	// mock data for testing
// 	tests := []struct {
// 		name    string
// 		input   map[string]interface{}
// 		want    *mysql.User
// 		wantErr *exception.HandleError
// 	}{
// 		{
// 			name: "Test User 1 Update Action",
// 			input: map[string]interface{}{
// 				"id":       strconv.FormatUint(uint64(result.ID), 10),
// 				"name":     "Test 123",
// 				"email":    "tes123@gmail.com",
// 				"password": "123456789",
// 			},
// 			want: &mysql.User{
// 				ID:       result.ID,
// 				Name:     "Test 123",
// 				Email:    "tes123@gmail.com",
// 				Password: "123456789",
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "Test User 2 Update Action (Email already exist)",
// 			input: map[string]interface{}{
// 				"id":       strconv.FormatUint(uint64(result.ID), 10),
// 				"name":     "Test 123",
// 				"email":    "tes123@gmail.com",
// 				"password": "123456789",
// 			},
// 			want: nil,
// 			wantErr: &exception.HandleError{
// 				Message:    utils.EmailAlreadyExist,
// 				Data:       map[string]interface{}{"email": "tes123@gmail.com"},
// 				StatusCode: http.StatusUnprocessableEntity,
// 			},
// 		},
// 	}

// 	// run test
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result, err := userAction.UserUpdateAction(ctx, test.input)
// 			if err != nil {
// 				assert.Equal(t, err.Message, test.wantErr.Message, fmt.Sprintf("got %v, want %v", err.Message, test.wantErr.Message))
// 				assert.Equal(t, err.StatusCode, test.wantErr.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, test.wantErr.StatusCode))
// 			}
// 			if result != nil {
// 				assert.Equal(t, result.Name, test.want.Name, fmt.Sprintf("got %v, want %v", result.Name, test.want.Name))
// 				assert.Equal(t, result.Email, test.want.Email, fmt.Sprintf("got %v, want %v", result.Email, test.want.Email))
// 				assert.Equal(t, result.Password, test.want.Password, fmt.Sprintf("got %v, want %v", result.Password, test.want.Password))
// 			}
// 		})
// 	}

// 	// Delete data after testing
// 	err := userAction.UserDeleteAction(ctx, request.UserGetOneModel{ID: result.ID})
// 	if err != nil {
// 		t.Errorf("UserDeleteAction() error = %v", err)
// 	}
// }
