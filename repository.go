package main

type BookRepository interface {
	List() ([]Book, error)
}

type LocalBookRepository struct{}

func (l LocalBookRepository) List() ([]Book, error) {
	return []Book{}, nil
}
