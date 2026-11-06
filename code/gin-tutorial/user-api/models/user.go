package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null;->:false" json:"-"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Role     string `gorm:"default:user" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)) == nil
}