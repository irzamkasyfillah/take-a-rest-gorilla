package request

type UserUpdateRequest struct {
	ID       uint   `json:"id" validate:"required,numeric"`
	Name     string `json:"name" validate:"omitempty,ascii"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty"`
}
