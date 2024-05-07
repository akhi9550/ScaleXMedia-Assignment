package models

type AddBookRequest struct {
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}

type BookDetails struct {
	BookName string `json:"book_name"`
}
