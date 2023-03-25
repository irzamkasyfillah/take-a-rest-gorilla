package action

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/exception"
	"github.com/irzam/my-app/api/user/utils"
	"gorm.io/gorm"
)

func (action *UserAction) UserUpdateAction(ctx context.Context, input map[string]interface{}) (*mysql.User, *exception.HandleError) {
	// Check id exist
	user_before, _ := action.UserRepository.GetByID(ctx, action.DB, input["id"].(uint))
	if user_before == nil || user_before.ID == 0 {
		return nil, &exception.HandleError{
			Message:    utils.UserNotFound,
			Data:       map[string]interface{}{"id": input["id"].(uint)},
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	// Check email exist
	if input["email"] != nil {
		if user, _ := action.UserRepository.GetByEmail(ctx, action.DB, input["email"].(string)); user != nil && user.ID != 0 {
			return nil, &exception.HandleError{
				Message:    utils.EmailAlreadyExist,
				Data:       map[string]interface{}{"email": input["email"]},
				StatusCode: http.StatusUnprocessableEntity,
			}
		}
	}

	// DB Transaction
	// child, cancel := context.WithTimeout(ctx, 1*time.Second)
	user, err := UserUpdateWithSnapshot(ctx, action, user_before, input)
	// cancel()
	if err != nil {
		return nil, &exception.HandleError{
			Message:    utils.Failed,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return user, nil
}

func UserUpdateWithSnapshot(ctx context.Context, action *UserAction, before *mysql.User, input map[string]interface{}) (after *mysql.User, er error) {
	db := action.DB.WithContext(ctx)
	er = db.Transaction(func(tx *gorm.DB) (err error) {
		// Update user
		after, err = action.UserRepository.Update(ctx, tx, input)
		if err != nil {
			return err
		}
		data, _ := json.Marshal(struct {
			Before *mysql.User `json:"before"`
			After  *mysql.User `json:"after"`
		}{before, after})
		// Create user history
		_, err = action.UserHistoryRepository.Create(ctx, tx, &mysql.UserHistory{
			UserID: before.ID,
			Action: "update",
			Data:   string(data),
		})
		if err != nil {
			return err
		}
		// time.Sleep(4 * time.Second)
		return nil
	})
	if er != nil {
		return nil, er
	}
	return after, nil
}
