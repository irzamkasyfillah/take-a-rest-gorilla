package request

type UserGetHistoryRequest struct {
	UserId      uint `json:"user_id" validate:"required,numeric"`
	PerPage     uint `json:"per_page" validate:"numeric"`
	CurrentPage uint `json:"current_page" validate:"numeric"`
}
