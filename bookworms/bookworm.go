package main

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
)

//loadBookworms reads the file and returns the list of bookworms

type Bookworm struct {
	Name string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buffedReader := bufio.NewReaderSize(f, 1024*1024)
	var bookworms []Bookworm

	err = json.NewDecoder(buffedReader).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

func booksCount(bookworms []Bookworm) map[Book] uint {
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

type byAuthor []Book

func (b byAuthor) Len() int { return len(b) }

func (b byAuthor) Swap(i, j int)  {
	b[i], b[j] = b[j], b[i]
}

func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
// 	sb := make(bookRecommendations)

// 	for _, bookworm := range bookworms {
// 		for i, book := range bookworm.Books {
// 			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
// 			registerBookRecommendations(sb, book, otherBooksOnShelves)
// 		}
// 	}

// 	recommendations := make([]Bookworm, len(bookworms))
// 	for i, bookworm := range bookworms {
// 		recommendations[i] = Bookworm{
// 			Name: bookworm.Name,
// 			Books: recommendBooks(sb, bookworm.Books),
// 		}
// 	}

// 	return recommendations
// }

type set map[Book]struct{}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}