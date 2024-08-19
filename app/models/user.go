package models

type User struct {
	ID       uint   `json:"user_id"`
	User     string `json:"-"`
	Name     string `json:"-"`
	Password string `json:"-"`
	UserType uint   `json:"user_type"`
}
