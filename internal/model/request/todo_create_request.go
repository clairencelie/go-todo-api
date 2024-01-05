package request

type TodoCreateRequest struct {
	UserId      int `json:"user_id"`
	Title       string
	Description string
}
