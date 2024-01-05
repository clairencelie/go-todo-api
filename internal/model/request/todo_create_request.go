package request

type TodoCreateRequest struct {
	UserId      int    `json:"user_id" validate:"required"`
	Title       string `validate:"required"`
	Description string
}
