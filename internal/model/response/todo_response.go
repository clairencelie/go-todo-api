package response

type TodoResponse struct {
	Id          int
	UserId      int
	Title       string
	Description string
	IsDone      bool
	CreatedAt   string
	UpdatedAt   string
}
