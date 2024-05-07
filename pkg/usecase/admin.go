package usecase

import (
	"errors"
	adminHelperInterface "scalexmedia-assignment/pkg/helper/interface"
	interfaces "scalexmedia-assignment/pkg/repository/interface"
	services "scalexmedia-assignment/pkg/usecase/interface"
	"scalexmedia-assignment/pkg/utils/models"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	adminHelper     adminHelperInterface.HelperAdmin
}

func NewAdminUseCase(repository interfaces.AdminRepository, helper adminHelperInterface.HelperAdmin) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repository,
		adminHelper:     helper,
	}
}

func (ad *adminUseCase) AdminLogin(email, name string) (*models.TokenAdmin, error) {
	details := models.AdminDetailsResponse{
		Email: email,
		Name:  name,
	}
	accessToken, err := ad.adminHelper.GenerateAccessTokenAdmin(details)
	if err != nil {
		return nil, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := ad.adminHelper.GenerateRefreshTokenAdmin(details)
	if err != nil {
		return nil, errors.New("counldn't create refreshtoken due to internal error")
	}

	return &models.TokenAdmin{
		Admin:        details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (ad *adminUseCase) ShowBooks(page, pageSize int) ([]models.BookDetails, error) {
	bookDetails, err := ad.adminRepository.ShowBooks(page, pageSize)
	if err != nil {
		return []models.BookDetails{}, err
	}
	return bookDetails, nil
}

func (ad *adminUseCase) AddBook(addBook models.AddBookRequest) error {
	bookExists := ad.adminRepository.CheckBookByName(addBook.BookName)
	if bookExists {
		return errors.New("book already exists")
	}
	err := ad.adminRepository.AddBook(addBook)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUseCase) DeleteBook(bookName models.BookDetails) error {
	bookExist := ad.adminRepository.CheckBookByName(bookName.BookName)
	if !bookExist {
		return errors.New("book doesn't exist")
	}
	err := ad.adminRepository.DeleteBook(bookName.BookName)
	if err != nil {
		return err
	}
	return nil
}
