package response

type StandardResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
