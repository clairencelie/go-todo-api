package request

type TodoUpdateRequest struct {
	Id          int    `validate:"required"`
	Title       string `validate:"required"`
	Description string
	IsDone      bool `json:"is_done"`
}
