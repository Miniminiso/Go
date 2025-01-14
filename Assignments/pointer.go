package Assignments

import "fmt"

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func CreateBook(b *Book) *Book {
	return b
}

func NewBook() {
	newBook1 := Book{Title: "Outlook", Author: "MS", Pages: 200, CopiesAvailable: 9}
	CreateBook(&newBook1)
	newBook1.Printbookdetails()
	newBook2 := Book{Title: "look", Author: "GO", Pages: 240, CopiesAvailable: 0}
	CreateBook(&newBook2)
	newBook2.Printbookdetails()
	fmt.Println(newBook1.Borrow())
	fmt.Println(newBook2.Borrow())
	newBook1.ReturnBook()
	newBook2.ReturnBook()
	SwapTitles(&newBook1, &newBook2)
}

func (b *Book) Printbookdetails() {
	fmt.Println(b.Title, b.Author, b.Pages, b.CopiesAvailable)
}

func (b *Book) Borrow() string {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable--
		return "Sucessful"
	} else {
		return "Not Available"
	}
}

func (b *Book) ReturnBook() {
	b.CopiesAvailable++
	fmt.Println(b.Title, b.Author, b.Pages, b.CopiesAvailable)
}

func SwapTitles(b1, b2 *Book) {
	b1.Title, b2.Title = b2.Title, b1.Title
	fmt.Println("Outlook Swaped to Look: ", b1.Title, b1.Author, b1.Pages, b1.CopiesAvailable)
	fmt.Println("look swaped to outlook: ", b2.Title, b2.Author, b2.Pages, b2.CopiesAvailable)
}
