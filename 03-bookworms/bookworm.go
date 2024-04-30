package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// because its a json type , we start the property name with uppercase letters
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworm(filepath string) ([]Bookworm, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	bookOnShelvs := booksCount(bookworms)
	var commonBooks []Book

	for book, count := range bookOnShelvs {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return commonBooks
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
