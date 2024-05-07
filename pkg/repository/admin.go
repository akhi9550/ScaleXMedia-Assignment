package repository

import (
	interfaces "scalexmedia-assignment/pkg/repository/interface"
	"scalexmedia-assignment/pkg/utils/models"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) ShowBooks(page, pageSize int) ([]models.BookDetails, error) {
	var bookDetails []models.BookDetails
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	err := ad.DB.Raw("SELECT book_name FROM books limit ? offset ?", pageSize, offset).Scan(&bookDetails).Error
	if err != nil {
		return []models.BookDetails{}, err
	}
	return bookDetails, nil
}

func (ad *adminRepository) CheckBookByName(bookName string) bool {
	var count int
	err := ad.DB.Raw(`SELECT COUNT(*) FROM books WHERE book_name ILIKE ?`, bookName).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (ad *adminRepository) AddBook(addBook models.AddBookRequest) error {
	err := ad.DB.Exec(`INSERT INTO books(book_name,author,publication_year) VALUES (?,?,?)`, addBook.BookName, addBook.Author, addBook.PublicationYear).Error
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) DeleteBook(bookName string) error {
	err := ad.DB.Exec(`DELETE FROM books WHERE book_name ILIKE ?`, bookName).Error
	if err != nil {
		return err
	}
	return nil
}
