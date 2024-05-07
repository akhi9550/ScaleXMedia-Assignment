package interfaces

type HelperAuth interface {
	GetTokenFromHeader(header string) string
	ExtractUserDetailsFromToken(tokenString string) (string, string, error)
	ExtractAdminFromToken(tokenString string) (string, string, error)
}
