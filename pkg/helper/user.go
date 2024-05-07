package helper

import (
	"scalexmedia-assignment/pkg/config"
	interfaces "scalexmedia-assignment/pkg/helper/interface"
	"scalexmedia-assignment/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type userHelper struct {
	cfg config.Config
}

func NewUserHelper(config config.Config) interfaces.HelperUser {
	return &userHelper{
		cfg: config,
	}
}

type UserClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h *userHelper) GenerateAccessTokenUser(user models.UserDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := h.GenerateToken(user.Name, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *userHelper) GenerateRefreshTokenUser(user models.UserDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := h.GenerateToken(user.Name, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *userHelper) GenerateToken(name, email string, expirationTime time.Time) (string, error) {
	claims := &AuthUserClaims{
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.KEY_USER))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
