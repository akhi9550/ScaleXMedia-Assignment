package helper

import (
	"scalexmedia-assignment/pkg/config"
	interfaces "scalexmedia-assignment/pkg/helper/interface"
	"scalexmedia-assignment/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type adminHelper struct {
	cfg config.Config
}

func NewAdminHelper(config config.Config) interfaces.HelperAdmin {
	return &adminHelper{
		cfg: config,
	}
}

type AdminClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func (h *adminHelper) GenerateAccessTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := h.GenerateAdminToken(admin.Email, admin.Name, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *adminHelper) GenerateRefreshTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := h.GenerateAdminToken(admin.Email, admin.Name, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *adminHelper) GenerateAdminToken(Email, Name string, expirationTime time.Time) (string, error) {
	claims := &AuthAdminClaims{
		Email: Email,
		Name:  Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.cfg.KEY_ADMIN))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
