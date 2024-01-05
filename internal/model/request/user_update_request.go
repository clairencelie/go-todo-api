package request

type UserUpdateRequest struct {
	Id          int
	Username    string
	Name        string
	Email       string
	PhoneNumber string `json:"phone_number"`
}
