package action

import (
	"context"
	"net/http"

	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/utils"
)

func (action *UserAction) UserGetAllAction(ctx context.Context, input *request.UserGetAllRequest) (interface{}, *exception.HandleError) {
	currentPage := uint(1)
	perPage := uint(10)
	if input.PerPage != 0 {
		perPage = input.PerPage
	}
	if input.CurrentPage != 0 {
		currentPage = input.CurrentPage
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
