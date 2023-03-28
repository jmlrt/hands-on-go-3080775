// challenges/types-composite/begin/main.go
package main

import "fmt"

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author author
}

// define a library type to track a list of books
// type library map[string][]book
type library struct {
	books []book
}

// define addBook to add a book to the library
func (l *library) addBook(book book) {
	l.books = append(l.books, book)
}

// define a lookupByAuthor function to find books by an author's name
func (l *library) lookupByAuthor(authorName string) []book {
	books := make([]book, 0)
	for _, b := range(l.books) {
		if b.author.name == authorName {
			books = append(books, b)
		}
	}
	return books
}

func main() {
	// create a new library
	myLibrary := library{}

	// add 2 books to the library by the same author
	myLibrary.addBook(book{"20,000 Leagues Under the Sea", author{"Jules Verne"}})
	myLibrary.addBook(book{"Around the World in 80 Days", author{"Jules Verne"}})

	// add 1 book to the library by a different author
	myLibrary.addBook(book{"Gulliver's Travels", author{"Jonathan Swift"}})

	// dump the library with spew
	fmt.Printf("%v\n", myLibrary)

	// lookup books by known author in the library
	lookupBooks := myLibrary.lookupByAuthor("Jules Verne")
	fmt.Printf("%v\n", lookupBooks)

	// print out the first book's title and its author's name
	firstBook := lookupBooks[0]
	fmt.Printf("%s - %s\n", firstBook.title, firstBook.author.name)

}
