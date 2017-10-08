package main

type Author struct {
	Name string
}

type Book struct {
	Id          int32
	Name        string
	Description string
	Pages       int32
	Authors     []Author
}

// sample authors
var author = Author{
	Name: "Tim Woo",
}

var me = Author{
	Name: "James Hrisho",
}

// sample books
var book = Book{
	Id:          100,
	Name:        "The Master Switch",
	Description: "A book about the internet",
	Pages:       200,
	Authors:     []Author{author},
}

var book2 = Book{
	Id:          101,
	Name:        "Some other book",
	Description: "A book not about the internet",
	Pages:       200,
	Authors:     []Author{author, me},
}

func findOne() Book {
	return book
}

func findAll() []Book {
	return []Book{book, book2}
}
