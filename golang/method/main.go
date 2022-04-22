package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

type Tbook struct {
	pages int
}

func (b *Tbook) SetPages(pages int) {
	b.pages = pages
}

func main() {
	var b = &Tbook{}
	b.SetPages(123)
	fmt.Println(b.pages)

	var book Book
	fmt.Printf("%T \n", book.Pages)
	fmt.Printf("%T \n", (&book).SetPages)

	var book1 = &Book{}
	book1.SetPages(2)
	book1.Pages()

	var book2 = Book{}
	book2.SetPages(2)
	book2.Pages()
}
