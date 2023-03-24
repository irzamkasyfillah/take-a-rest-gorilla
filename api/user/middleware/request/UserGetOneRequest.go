package request

type UserGetOneRequest struct {
	ID uint `validate:"required,numeric"`
}
