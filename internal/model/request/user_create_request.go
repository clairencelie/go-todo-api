package request

type UserCreateRequest struct {
	Username    string `validate:"required"`
	Password    string `validate:"required"`
	Name        string `validate:"required"`
	Email       string `validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
