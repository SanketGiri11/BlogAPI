package models

type Tags struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Tags   string `json:"tags"`
	PostID int    `json:"postid"`
	Posts  Posts  `json:"user" gorm:"foreignkey:PostID"`
}
