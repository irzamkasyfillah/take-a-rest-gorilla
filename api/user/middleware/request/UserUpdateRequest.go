package request

type UserUpdateRequest struct {
	ID       uint   `json:"id" validate:"required,numeric"`
	Name     string `json:"name" validate:"alpha"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:""`
}
