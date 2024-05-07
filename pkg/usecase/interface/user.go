package interfaces

import "scalexmedia-assignment/pkg/utils/models"

type UserUseCase interface {
	UserLogin(email,name string) (*models.TokenUser, error)
}
