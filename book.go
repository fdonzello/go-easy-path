package main

import "fmt"

type Book struct {
	Name string
}

func (b Book) Read() {
	fmt.Println("reading ", b.Name)
}
