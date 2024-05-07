package interfaces

import (
	"scalexmedia-assignment/pkg/utils/models"
)

type AdminRepository interface {
	ShowBooks(page, pageSize int) ([]models.BookDetails, error)
	CheckBookByName(bookName string) bool
	AddBook(addBook models.AddBookRequest) error
	DeleteBook(bookName string) error
}
