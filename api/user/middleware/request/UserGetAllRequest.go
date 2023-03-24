package request

type UserGetAllRequest struct {
	PerPage     uint `json:"per_page" validate:"numeric"`
	CurrentPage uint `json:"current_page" validate:"numeric"`
}
