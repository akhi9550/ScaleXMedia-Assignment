package models

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"min=6,max=20"`
}

type AdminDetailsResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TokenAdmin struct {
	Admin        AdminDetailsResponse
	AccessToken  string
	RefreshToken string
}

type AdminStruct struct {
	Name     string
	Password string
}
