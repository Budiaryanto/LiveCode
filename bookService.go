package main

import (
	"fmt"
	"strings"
)

type BookService struct {
	r BookRepository
}

func (bs *BookService) registerNewBook(b Buku) {
	bs.r.AddNewBook(b)
}

func (bs *BookService) printAllBookCollection() {
	collection := bs.r.FindAllBook()
	bs.printConsole("Koleksi Buku", collection)
}

func (bs *BookService) searchBookById(id string) {
	var funcFilter = func(b Buku) bool {
		return b.id == id
	}
	result := bs.r.FindBookByField(funcFilter)
	bs.printConsole(fmt.Sprintf("Hasil Pencarian Buku dengan Id : %s", id), result)
}

func (bs *BookService) searchBookByTitle(title string) {
	var funcFilter = func(b Buku) bool {
		return strings.Contains(b.nama, title)
	}
	result := bs.r.FindBookByField(funcFilter)
	bs.printConsole(fmt.Sprintf("Hasil Pencarian Buku dengan kata Kunci : %s", title), result)
}
func (bs *BookService) printConsole(header string, books []Buku) {
	fmt.Println("")
	fmt.Println(header)
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("%-10s%-30s%-15s%-15s%10s\n", "ID", "Title", "Author", "Genre", "Price")
	fmt.Printf("%s\n", strings.Repeat("-", 80))
	for _, b := range books {
		fmt.Printf("%-10s%-30s%-15s%-15s%10.2f\n", b.id, b.nama, b.pengarang, b.genre, b.harga)
	}
}
