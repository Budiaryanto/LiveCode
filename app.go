package main

func main() {
	book01 := Buku{
		id:        "123",
		nama:      "Samson Betawi",
		pengarang: "Gilbert",
		genre:     "Superhero",
		harga:     12500,
	}
	book02 := Buku{
		id:        "12356",
		nama:      "jhonson",
		pengarang: "Aliba",
		genre:     "Majalah",
		harga:     12000,
	}
	book03 := Buku{
		id:        "800 - 900",
		nama:      "Dragon Ball Super 3",
		pengarang: "Gramedia",
		genre:     "Komik",
		harga:     10000,
	}
	repo := BookRepository{}
	bookService := BookService{r: repo}

	bookService.registerNewBook(book01)
	bookService.registerNewBook(book02)
	bookService.registerNewBook(book03)
	bookService.printAllBookCollection()

	bookService.searchBookByTitle("Dragon")
	bookService.searchBookById("123")
}
