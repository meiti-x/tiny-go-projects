package main

import "fmt"

func main() {
	bookmars, err := loadBookworm("testdata/03-bookworms.json")

	if err != nil {
		panic(err)
	}
	fmt.Println(bookmars)
}
