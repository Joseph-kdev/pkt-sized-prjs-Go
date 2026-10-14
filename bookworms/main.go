package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "testdata/bookworms.json", "The path of the booworms json file...")

	bookworms, err := loadBookworms(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}