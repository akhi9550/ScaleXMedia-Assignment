package interfaces

import (
	"scalexmedia-assignment/pkg/utils/models"
	"time"
)

type HelperAdmin interface {
	GenerateAccessTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	GenerateRefreshTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	GenerateAdminToken(Email, Name string, expirationTime time.Time) (string, error)
}
