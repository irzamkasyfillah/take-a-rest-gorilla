package __test__

import (
	"context"
	"fmt"
	"testing"

	"github.com/irzam/my-app/api/user/exception"
	"github.com/stretchr/testify/assert"
)

// unit test without connection to database
func TestUserGetAllActionMock(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]interface{}
		wantErr *exception.HandleError
	}{
		{
			name: "Test User Get All Action",
			input: map[string]interface{}{
				"current_page": float64(1),
				"per_page":     float64(10),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserRepository.Mock.On("GetAll", context.Background(), 1, 10).Return(nil, nil)
			_, err := UserAction.UserGetAllAction(context.Background(), tt.input)
			if err != nil {
				assert.Equal(t, err, tt.wantErr, fmt.Sprintf("got %v, want %v", err, tt.wantErr))
			}
		})
		UserRepository.Mock.ExpectedCalls = nil
	}
}

// unit test + connect db
// func TestUserGetAllAction(t *testing.T) {
// 	userHistoryRepository := repository.NewUserHistoryRepository()
// 	userRepository := repository.NewUserRepository()
// 	userAction := action.NewUserAction(userRepository, userHistoryRepository, SetupDB())

// 	// Create a new context
// 	ctx := context.Background()

// 	tests := []struct {
// 		name    string
// 		input   map[string]interface{}
// 		wantErr *exception.HandleError
// 	}{
// 		{
// 			name: "Test User Get All Action",
// 			input: map[string]interface{}{
// 				"current_page": float64(1),
// 				"per_page":     float64(10),
// 			},
// 			wantErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := userAction.UserGetAllAction(ctx, tt.input)
// 			if err != nil {
// 				assert.Equal(t, err, tt.wantErr, fmt.Sprintf("got %v, want %v", err, tt.wantErr))
// 			}
// 		})
// 	}
// }
