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
	"github.com/stretchr/testify/mock"
)

// unit test without connection to database
func TestUserDeleteActionMock(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]interface{}
		userExist *mysql.User
		wantErr   *exception.HandleError
	}{
		{
			name: "Test User Delete Action",
			input: map[string]interface{}{
				"id": uint(1),
			},
			userExist: &mysql.User{
				ID:       1,
				Name:     "Test",
				Email:    "tes@gmail.com",
				Password: "123456",
			},
			wantErr: nil,
		},
		{
			name: "Test User Delete Action (User not found)",
			input: map[string]interface{}{
				"id": uint(1),
			},
			userExist: nil,
			wantErr: &exception.HandleError{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    utils.UserNotFound,
				Data: map[string]interface{}{
					"id": uint(1),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserHistoryRepository.Mock.On("Create", context.Background(), mock.Anything).Return(nil)
			UserRepository.Mock.On("GetByID", context.Background(), tt.input["id"].(uint)).Return(tt.userExist, nil)
			UserRepository.Mock.On("Delete", context.Background(), tt.input["id"].(uint)).Return(nil)

			err := UserAction.UserDeleteAction(context.Background(), request.UserGetOneModel{ID: tt.input["id"].(uint)})
			if err != nil {
				assert.Equal(t, err.StatusCode, tt.wantErr.StatusCode, fmt.Sprintf("got %v, want %v", err.StatusCode, tt.wantErr.StatusCode))
				assert.Equal(t, err.Message, tt.wantErr.Message, fmt.Sprintf("got %v, want %v", err.Message, tt.wantErr.Message))
			}
		})
		UserRepository.Mock.ExpectedCalls = nil
	}
}
