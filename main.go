package main

import (
	"fmt"
	"log"
)

func main() {
	repo := LocalBookRepository{}
	getAllBooks(repo)
}

func getAllBooks(r BookRepository) {
	books, err := r.List()

	if err != nil {
		log.Fatalf("error on retrieving books: %v", err)
	}

	fmt.Println(books)
}
