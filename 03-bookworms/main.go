package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filepath := flag.String("filepath", "testdata/bookworms.json", "declare bookworms json file path")
	flag.Parse()
	bookworms, err := loadBookworm(*filepath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)

}
