package models

type Posts struct {
	Id     uint   `json:"id"`
	Titel  string `json:"titel"`
	Desc   string `json:"Desc"`
	UserID int    `json:"userid"`
	User   User   `json:"user" gorm:"foreignkey:UserID"`
}
