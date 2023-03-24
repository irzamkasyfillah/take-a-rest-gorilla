package mock

import (
	"context"

	"github.com/irzam/my-app/api/user/entity/model/mysql"
	"github.com/irzam/my-app/api/user/middleware/request"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

type UserRepositoryMockInterface interface {
	GetByEmail(ctx context.Context, db *gorm.DB, email string) (*mysql.User, error)
	GetByID(ctx context.Context, db *gorm.DB, id uint) (*mysql.User, error)
	GetAll(ctx context.Context, db *gorm.DB, args ...uint) (interface{}, error)
	Create(ctx context.Context, db *gorm.DB, input *request.UserCreateRequest) (*mysql.User, error)
	Update(ctx context.Context, db *gorm.DB, input map[string]interface{}) (*mysql.User, error)
	Delete(ctx context.Context, db *gorm.DB, id uint) error
}

func (u *UserRepositoryMock) GetByID(ctx context.Context, db *gorm.DB, id uint) (*mysql.User, error) {
	args := u.Mock.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*mysql.User), args.Error(1)
	}
}

func (u *UserRepositoryMock) GetByEmail(ctx context.Context, db *gorm.DB, email string) (*mysql.User, error) {
	args := u.Mock.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*mysql.User), args.Error(1)
	}
}

func (u *UserRepositoryMock) GetAll(ctx context.Context, db *gorm.DB, args ...uint) (interface{}, error) {
	var currentPage, perPage uint
	if len(args) > 0 {
		currentPage = args[0]
		perPage = args[1]
	} else {
		currentPage = 1
		perPage = 10
	}
	arg := u.Mock.Called(ctx, currentPage, perPage)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		return arg.Get(0), arg.Error(1)
	}
	arg = u.Mock.Called(ctx)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		return arg.Get(0), arg.Error(1)
	}
}

func (u *UserRepositoryMock) Create(ctx context.Context, db *gorm.DB, input *request.UserCreateRequest) (*mysql.User, error) {
	// child, _ := context.WithTimeout(ctx, 2*time.Second)
	args := u.Mock.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*mysql.User), nil
	}
}

func (u *UserRepositoryMock) Update(ctx context.Context, db *gorm.DB, input map[string]interface{}) (*mysql.User, error) {
	args := u.Mock.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*mysql.User), nil
	}
}

func (u *UserRepositoryMock) Delete(ctx context.Context, db *gorm.DB, id uint) error {
	args := u.Mock.Called(ctx, id)
	return args.Error(0)
}
