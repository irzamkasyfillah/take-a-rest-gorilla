package action

import (
	"context"
	"net/http"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/utils"
)

func (action *UserAction) UserGetOneAction(ctx context.Context, input request.UserGetOneModel) (*mysql.User, *exception.HandleError) {
	user, _ := action.UserRepository.GetByID(ctx, action.DB, input.ID)
	if user == nil || user.ID == 0 {
		return nil, &exception.HandleError{
			Message:    utils.UserNotFound,
			Data:       map[string]interface{}{"id": input.ID},
			StatusCode: http.StatusUnprocessableEntity,
		}
	}
	return user, nil
}
