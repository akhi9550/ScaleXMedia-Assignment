package interfaces

import "scalexmedia-assignment/pkg/utils/models"

type AdminUseCase interface {
	AdminLogin(email, name string) (*models.TokenAdmin, error)
	ShowBooks(page, pageSize int) ([]models.BookDetails, error)
	AddBook(addBook models.AddBookRequest) error
	DeleteBook(bookName models.BookDetails) error
}
