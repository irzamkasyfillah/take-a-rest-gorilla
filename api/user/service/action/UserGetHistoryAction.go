package action

import (
	"context"
	"net/http"

	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/utils"
)

func (action *UserAction) UserGetHistoryAction(ctx context.Context, payload map[string]interface{}) (interface{}, *exception.HandleError) {
	// Check if payload user_id is exist
	if payload["user_id"] == nil {
		return nil, &exception.HandleError{
			Data:       map[string]interface{}{"message": "User id is required"},
			Message:    utils.UserNotFound,
			StatusCode: http.StatusNotFound,
		}
	}

	currentPage := 1
	perPage := 10
	if payload["per_page"] != nil {
		perPage = int(payload["per_page"].(float64))
	}
	if payload["current_page"] != nil {
		currentPage = int(payload["current_page"].(float64))
	}

	// Get all user
	users, err := action.UserHistoryRepository.GetByUserID(ctx, action.DB, uint(payload["user_id"].(float64)), currentPage, perPage)
	if err != nil {
		return nil, &exception.HandleError{
			Message:    utils.Failed,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return users, nil
}
