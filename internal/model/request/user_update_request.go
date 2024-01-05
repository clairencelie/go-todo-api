package request

type UserUpdateRequest struct {
	Id          int    `validate:"required"`
	Username    string `validate:"required"`
	Name        string `validate:"required"`
	Email       string `validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
