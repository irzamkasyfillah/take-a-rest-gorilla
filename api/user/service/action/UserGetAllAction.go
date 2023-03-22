package action

import (
	"context"
	"net/http"

	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/utils"
)

func (action *UserAction) UserGetAllAction(ctx context.Context, payload map[string]interface{}) (interface{}, *exception.HandleError) {
	currentPage := 1
	perPage := 10
	if payload["per_page"] != nil {
		perPage = int(payload["per_page"].(float64))
	}
	if payload["current_page"] != nil {
		currentPage = int(payload["current_page"].(float64))
	}

	// Get all user
	users, err := action.UserRepository.GetAll(ctx, action.DB, currentPage, perPage)
	if err != nil {
		return nil, &exception.HandleError{
			Message:    utils.Failed,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return users, nil
}
