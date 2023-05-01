package user

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Store struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewStore(log *zap.SugaredLogger, db *gorm.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

func (s Store) Create(ctx context.Context, nu NewUser, now time.Time) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: passwordHash,
		DateCreated:  now,
	}
	result := s.db.Create(&u)
	if result.Error != nil {
		return User{}, result.Error
	}
	return u, nil
}

func (s Store) QueryByEmail(ctx context.Context, email string) (User, error) {
	var usr User
	res := s.db.Where(&User{Email: email}).First(&usr)
	if res.Error != nil {
		return User{}, res.Error
	}
	return usr, nil

}
