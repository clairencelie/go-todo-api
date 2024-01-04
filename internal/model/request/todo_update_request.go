package request

type TodoUpdateRequest struct {
	Title       string
	Description string
	IsDone      bool
}
