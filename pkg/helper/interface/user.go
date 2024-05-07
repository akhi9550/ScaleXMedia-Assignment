package interfaces

import (
	"scalexmedia-assignment/pkg/utils/models"
	"time"
)

type HelperUser interface {
	GenerateAccessTokenUser(user models.UserDetailsResponse) (string, error)
	GenerateRefreshTokenUser(user models.UserDetailsResponse) (string, error)
	GenerateToken(name, email string, expirationTime time.Time) (string, error)
}
