package swagger

import "encoding/json"

var books = map[string]Book{
	"12-34567-89-0": Book{Isbn: "12-34567-89-0", Title: "Exemplo de Livro", Author: "Rodrigo Moura", Edition: "1"},
}

// BookToJSON to be used for marshalling of Book type
func (b Book) BookToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// BookFromJSON to be used for unmarshalling of Book type
func BookFromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.Isbn]
	if exists {
		return "", false
	}
	books[book.Isbn] = book
	return book.Isbn, true
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}
