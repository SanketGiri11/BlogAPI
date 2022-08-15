package models

type Comments struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Comments string `json:"comments"`
	PostID   int    `json:"postid"`
	Posts    Posts  `json:"user" gorm:"foreignkey:PostID"`
}
