package user

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Roles        pq.StringArray `gorm:"type:text[]" json:"roles"`
	PasswordHash []byte         `json:"-"`
	DateCreated  time.Time      `json:"date_created"`
	DateUpdated  *time.Time     `json:"date_updated"`
}

func (u User) IsValidPassword(pwd string) bool {
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(pwd)); err != nil {
		return false
	}
	return true
}

type NewUser struct {
	Name            string   `json:"name" binding:"required"`
	Email           string   `json:"email" binding:"required,email"`
	Roles           []string `json:"roles" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	PasswordConfirm string   `json:"password_confirm" binding:"eqfield=Password"`
}
