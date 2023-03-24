package request

type UserDeleteRequest struct {
	ID uint `validate:"required,numeric"`
}
