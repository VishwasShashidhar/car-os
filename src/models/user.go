package models

type User struct {
	Base
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}
