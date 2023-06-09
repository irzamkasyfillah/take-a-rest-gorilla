package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,ascii"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
