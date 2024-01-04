package request

type TodoUpdateRequest struct {
	Id          int
	Title       string
	Description string
	IsDone      bool
}
