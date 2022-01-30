package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalf("please provide a source: remote or local")
	}

	source := args[0]

	var r BookRepository

	switch source {
	case "remote":
		fmt.Println("retrieving books from remote")
		r = RemoteBookRepository{}
	case "local":
		fmt.Println("retrieving books from local")
		r = LocalBookRepository{}
	default:
		log.Fatalf("unknown source, available are: remote, local")
	}

	getAllBooks(r)
}

func getAllBooks(r BookRepository) {
	books, err := r.List()

	if err != nil {
		log.Fatalf("error on retrieving books: %v", err)
	}

	fmt.Println(books)
}
