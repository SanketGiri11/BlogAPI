package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"id"`
	Author   string `json:"author"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Satus    string `json:"status"`
}

//HashPassword function converted the Password to a encrypted format for storing it in database
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//Checkpassword fuction compare the password with the password stored in DB
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
