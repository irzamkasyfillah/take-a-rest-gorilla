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

func (action *UserAction) UserCreateAction(ctx context.Context, input *mysql.User) (*mysql.User, *exception.HandleError) {
	// Check email exist
	if user, _ := action.UserRepository.GetByEmail(ctx, action.DB, input.Email); user != nil && user.ID != 0 {
		return nil, &exception.HandleError{
			Message:    utils.EmailAlreadyExist,
			Data:       map[string]interface{}{"email": input.Email},
			StatusCode: http.StatusUnprocessableEntity,
		}
	}
	// ============= DB Transaction =================
	/**
	 * Without using go routine and channel
	 */
	// user, err := UserCreateWithSnapshot(input)
	/**
	 * Example implement go routine and channel
	 */
	ch1 := make(chan *mysql.User)
	ch2 := make(chan error)
	// impl context with timeout 2 second
	// child, cancel := context.WithTimeout(ctx, 2*time.Second)
	go func() {
		user, err := UserCreateWithSnapshot(ctx, action, input)
		ch1 <- user
		ch2 <- err
	}()
	user := <-ch1
	err := <-ch2
	// cancel()
	// ============= DB Transaction =================
	if err != nil {
		return nil, &exception.HandleError{
			Message:    utils.Failed,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return user, nil
}

func UserCreateWithSnapshot(ctx context.Context, action *UserAction, input *mysql.User) (user *mysql.User, er error) {
	db := action.DB.WithContext(ctx)
	er = db.Transaction(func(tx *gorm.DB) (err error) {
		// Create user
		user, err = action.UserRepository.Create(ctx, tx, input)
		if err != nil {
			return err
		}
		data, _ := json.Marshal(struct {
			Before *mysql.User `json:"before"`
			After  *mysql.User `json:"after"`
		}{nil, user})
		// Create user history
		_, err = action.UserHistoryRepository.Create(ctx, tx, &mysql.UserHistory{
			UserID: user.ID,
			Action: "create",
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
	return user, nil
}

// func UserCreateWithSnapshot(ctx context.Context, db *gorm.DB, input *mysql.User) (*mysql.User, error) {
// 	db = db.WithContext(ctx)
// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		// Create user
// 		if err := tx.Table(mysql.UserTableName()).Create(input).Error; err != nil {
// 			return err
// 		}
// 		data, _ := json.Marshal(struct {
// 			Before *mysql.User `json:"before"`
// 			After  *mysql.User `json:"after"`
// 		}{nil, input})
// 		// Create user history
// 		if err := tx.Table(mysql.UserHistoryTableName()).Create(&mysql.UserHistory{
// 			UserID: input.ID,
// 			Action: "create",
// 			Data:   string(data),
// 		}).Error; err != nil {
// 			return err
// 		}
// 		// time.Sleep(4 * time.Second)
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return input, nil
// }
