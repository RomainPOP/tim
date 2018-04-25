package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name			string `gorm:"size:255"`
	Email 			string `gorm:"type:varchar(100);unique_index"`
	HashedPassword	[]byte
	Active 			bool
	isAdmin 		bool
}

// SetNewPassword set a new hashed password to user
func (user *User)  SetNewPassword(passwordString string){
	bcrytpassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	user.HashedPassword = bcrytpassword
}