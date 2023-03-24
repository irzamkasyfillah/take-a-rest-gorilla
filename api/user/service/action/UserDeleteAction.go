package action

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/irzam/my-app/api/user/utils"
	"gorm.io/gorm"
)

func (action *UserAction) UserDeleteAction(ctx context.Context, input *request.UserDeleteRequest) *exception.HandleError {
	// Check if user exist
	user, _ := action.UserRepository.GetByID(ctx, action.DB, input.ID)
	if user == nil || user.ID == 0 {
		return &exception.HandleError{
			Message:    utils.UserNotFound,
			Data:       map[string]interface{}{"id": input.ID},
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	// DB Transaction
	err := UserDeleteWithSnapshot(ctx, action, user)
	if err != nil {
		return &exception.HandleError{
			Message:    utils.Failed,
			Data:       map[string]interface{}{"id": input.ID},
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}

func UserDeleteWithSnapshot(ctx context.Context, action *UserAction, before *mysql.User) (er error) {
	db := action.DB.WithContext(ctx)
	er = db.Transaction(func(tx *gorm.DB) (err error) {
		// Delete user
		err = action.UserRepository.Delete(ctx, tx, before.ID)
		if err != nil {
			return err
		}
		data, _ := json.Marshal(struct {
			Before *mysql.User `json:"before"`
			After  *mysql.User `json:"after"`
		}{
			Before: before,
			After:  nil,
		})
		// Create history
		_, err = action.UserHistoryRepository.Create(ctx, tx, &mysql.UserHistory{
			UserID: before.ID,
			Action: "delete",
			Data:   string(data),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if er != nil {
		return er
	}
	return nil
}
