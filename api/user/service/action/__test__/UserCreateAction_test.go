package __test__

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/service/action"
	myMock "github.com/irzam/my-app/api/user/service/action/__test__/mock"
	"github.com/irzam/my-app/api/user/utils"
	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserHistoryRepository = &myMock.UserHistoryRepositoryMock{Mock: mock.Mock{}}
var UserRepository = &myMock.UserRepositoryMock{Mock: mock.Mock{}}
var UserAction = action.NewUserAction(UserRepository, UserHistoryRepository, SetupDB())

func SetupDB() (db *gorm.DB) {
	var err error
	// Connect to the database
	dsn := "root:@tcp(localhost:3306)/user"
	db, err = gorm.Open(sql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to database")
	return db
}

func TestMain(m *testing.M) {
	log.Println("⚠️⚠️ Unit Test for User Action is running ⚠️⚠️")
	m.Run()
	log.Println("✅✅ Unit Test for User Action is done ✅✅")
}

// unit test without connect db
func TestUserCreateActionMock(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		input      *mysql.User
		want       *mysql.User
		emailExist *mysql.User
		wantErr    *exception.HandleError
	}{
		{
			name: "Test User Create Action",
			input: &mysql.User{
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			emailExist: nil,
			want: &mysql.User{
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			wantErr: nil,
		},
		{
			name: "Test User Create Action (Email already exists)",
			input: &mysql.User{
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			emailExist: &mysql.User{
				ID:       1,
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			want: nil,
			wantErr: &exception.HandleError{
				Message:    utils.EmailAlreadyExist,
				Data:       map[string]interface{}{"email": "tes@gmail.com"},
				StatusCode: http.StatusUnprocessableEntity,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UserRepository.Mock.On("GetByEmail", ctx, test.input.Email).Return(test.emailExist, nil)
			UserRepository.Mock.On("Create", ctx, test.input).Return(test.want, test.wantErr)
			UserHistoryRepository.Mock.On("Create", ctx, mock.Anything).Return(nil, nil)

			result, err := UserAction.UserCreateAction(ctx, test.input)
			if result != nil {
				assert.Equal(t, test.want.Name, result.Name, fmt.Sprintf("got %v, want %v", result.Name, test.want.Name))
				assert.Equal(t, test.want.Email, result.Email, fmt.Sprintf("got %v, want %v", result.Email, test.want.Email))
				assert.Equal(t, test.want.Password, result.Password, fmt.Sprintf("got %v, want %v", result.Password, test.want.Password))
			}
			if err != nil {
				assert.Equal(t, test.wantErr.Message, err.Message, fmt.Sprintf("got %v, want %v", err.Message, test.wantErr.Message))
				assert.Equal(t, test.wantErr.StatusCode, err.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, test.wantErr.StatusCode))
			}
		})
		UserRepository.Mock.ExpectedCalls = nil
	}
}

// unit test + connect db
// func TestUserCreateAction(t *testing.T) {
// 	userHistoryRepository := repository.NewUserHistoryRepository()
// 	// userRepository := &myMock.UserRepositoryMock{Mock: mock.Mock{}}
// 	userRepository := repository.NewUserRepository()
// 	userAction := action.NewUserAction(userRepository, userHistoryRepository, SetupDB())

// 	ctx := context.Background()

// 	tests := []struct {
// 		name    string
// 		input   *mysql.User
// 		want    *mysql.User
// 		wantErr *exception.HandleError
// 	}{
// 		{
// 			name: "Test User 1 Create Action",
// 			input: &mysql.User{
// 				Name:     "Test",
// 				Email:    "tes@gmail.com",
// 				Password: "123456",
// 			},
// 			want: &mysql.User{
// 				Name:     "Test",
// 				Email:    "tes@gmail.com",
// 				Password: "123456",
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "Test User 2 Create Action (Email already exists)",
// 			input: &mysql.User{
// 				Name:     "Test",
// 				Email:    "tes@gmail.com",
// 				Password: "123456",
// 			},
// 			want: nil,
// 			wantErr: &exception.HandleError{
// 				Message:    utils.EmailAlreadyExist,
// 				Data:       map[string]interface{}{"email": "tes@gmail.com"},
// 				StatusCode: http.StatusUnprocessableEntity,
// 			},
// 		},
// 	}
// 	var id uint
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result, err := userAction.UserCreateAction(ctx, test.input)
// 			if err != nil {
// 				assert.Equal(t, test.wantErr.Message, err.Message, fmt.Sprintf("got %v, want %v", err.Message, test.wantErr.Message))
// 				assert.Equal(t, test.wantErr.StatusCode, err.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, test.wantErr.StatusCode))
// 			}
// 			if result != nil {
// 				id = result.ID
// 				assert.Equal(t, test.want.Name, result.Name, fmt.Sprintf("got %v, want %v", result.Name, test.want.Name))
// 				assert.Equal(t, test.want.Email, result.Email, fmt.Sprintf("got %v, want %v", result.Email, test.want.Email))
// 				assert.Equal(t, test.want.Password, result.Password, fmt.Sprintf("got %v, want %v", result.Password, test.want.Password))
// 			}
// 		})
// 	}
// 	err := userAction.UserDeleteAction(ctx, request.UserGetOneModel{ID: id})
// 	if err != nil {
// 		t.Errorf("UserDeleteAction() error = %v", err)
// 	}
// }
