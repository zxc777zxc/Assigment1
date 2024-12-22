package main

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Println("Book already exists")
		return
	}
	l.Books[book.ID] = book
	fmt.Println("Book added successfully")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed")
		return
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed successfully")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book is not borrowed")
		return
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned successfully")
}

func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available")
		return
	}
	fmt.Println("List of books:")
	for _, book := range l.Books {
		status := "Available"
		if book.IsBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("ID: %s, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
	}
}
