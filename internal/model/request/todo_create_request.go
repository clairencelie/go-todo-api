package request

type TodoCreateRequest struct {
	UserId      int
	Title       string
	Description string
}
