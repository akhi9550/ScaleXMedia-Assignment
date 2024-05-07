package models

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"min=6,max=20"`
}

type UserDetailsResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TokenUser struct {
	User         UserDetailsResponse
	AccessToken  string
	RefreshToken string
}
type UserStruct struct {
	Name     string
	Password string
}
