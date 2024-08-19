package models

type User struct {
	ID       uint   `json:"user_id"`
	Username     string `json:"-"`
	Name     string `json:"-"`
	Password string `json:"-"`
	UserType uint   `json:"user_type"`
}
