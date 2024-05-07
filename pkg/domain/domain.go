package domain

type Book struct {
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}
