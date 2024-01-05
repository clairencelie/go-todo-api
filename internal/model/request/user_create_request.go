package request

type UserCreateRequest struct {
	Username    string
	Password    string
	Name        string
	Email       string
	PhoneNumber string `json:"phone_number"`
}
