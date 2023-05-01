package user

import (
	"RssReader/bussiness/data/store/user"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Core struct {
	log  *zap.SugaredLogger
	user user.Store
}

func NewCore(log *zap.SugaredLogger, db *gorm.DB) Core {
	return Core{
		log:  log,
		user: user.NewStore(log, db),
	}
}

func (c Core) Create(ctx context.Context, nu user.NewUser, now time.Time) (user.User, error) {
	usr, err := c.user.Create(ctx, nu, now)
	if err != nil {
		return user.User{}, fmt.Errorf("create: %w", err)
	}
	return usr, nil
}

func (c Core) Authenticate(ctx context.Context, now time.Time, email, password string) (jwt.Claims, error) {
	usr, err := c.user.QueryByEmail(ctx, email)
	if !usr.IsValidPassword(password) {
		return jwt.StandardClaims{}, fmt.Errorf("invalid password")
	}
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(time.Minute * 100).Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
		Issuer:    "web service",
		Subject:   strconv.Itoa(usr.ID),
	}
	if err != nil {
		return jwt.StandardClaims{}, err
	}
	return claims, nil
}

//func (c Core) Update(ctx context.Context, claims auth.Claims, userID string, uu user.UpdateUser, now time.Time) error {
//	if err := c.user.Update(ctx, claims, userID, uu, now); err != nil {
//		return fmt.Errorf("update: %w", err)
//	}
//	return nil
//}
