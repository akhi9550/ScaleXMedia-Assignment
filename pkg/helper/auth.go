package helper

import (
	"fmt"
	"scalexmedia-assignment/pkg/config"
	interfaces "scalexmedia-assignment/pkg/helper/interface"

	"github.com/golang-jwt/jwt"
)

type authHelper struct {
	cfg config.Config
}

func NewAuthHelper(config config.Config) interfaces.HelperAuth {
	return &authHelper{
		cfg: config,
	}
}

type AuthUserClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type AuthAdminClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h *authHelper) GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}
func (h *authHelper) ExtractUserDetailsFromToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(h.cfg.KEY_USER), nil
	})

	if err != nil {
		fmt.Println("errors:-", err)
		return "", "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}

	return claims.Email, claims.Name, nil

}

func (h *authHelper) ExtractAdminFromToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthAdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(h.cfg.KEY_ADMIN), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*AuthAdminClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}

	return claims.Email, claims.Name, nil

}
