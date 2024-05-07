package usecase

import (
	interfaces "scalexmedia-assignment/pkg/helper/interface"
	services "scalexmedia-assignment/pkg/usecase/interface"
	"scalexmedia-assignment/pkg/utils/models"

	"errors"
)

type userUseCase struct {
	userHelper interfaces.HelperUser
}

func NewUserUseCase(helper interfaces.HelperUser) services.UserUseCase {
	return &userUseCase{
		userHelper: helper,
	}
}

func (ur *userUseCase) UserLogin(email, name string) (*models.TokenUser, error) {
	Details := models.UserDetailsResponse{
		Email: email,
		Name:  name,
	}
	accessToken, err := ur.userHelper.GenerateAccessTokenUser(Details)
	if err != nil {
		return nil, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := ur.userHelper.GenerateRefreshTokenUser(Details)
	if err != nil {
		return nil, errors.New("counldn't create refreshtoken due to internal error")
	}

	return &models.TokenUser{
		User:         Details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
