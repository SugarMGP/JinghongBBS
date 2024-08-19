package models

type Report struct {
	User    uint   `json:"-"`
	Post    uint   `json:"post_id"`
	Content string `json:"content"`
	Reason  string `json:"reason"`
	Status  uint   `json:"status"`
}
