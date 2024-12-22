package Library

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	library := Library{Books: make(map[string]Book)}
	reader := bufio.NewReader(os.Stdin)
	idCounter := 0

	for {
		fmt.Println("\nMenu:\n1. Add a book\n2. Borrow a book\n3. Return a book\n4. List available books\n5. Exit")
		fmt.Print("Enter your choice: ")
		choiceInput, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceInput))
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter book title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter book author: ")
			author, _ := reader.ReadString('\n')

			book := Book{
				ID:         strconv.Itoa(idCounter),
				Title:      strings.TrimSpace(title),
				Author:     strings.TrimSpace(author),
				IsBorrowed: false,
			}
			library.AddBook(book)
			idCounter++

		case 2:
			fmt.Print("Enter book ID to borrow: ")
			id, _ := reader.ReadString('\n')
			library.BorrowBook(strings.TrimSpace(id))

		case 3:
			fmt.Print("Enter book ID to return: ")
			id, _ := reader.ReadString('\n')
			library.ReturnBook(strings.TrimSpace(id))

		case 4:
			library.ListBooks()

		case 5:
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
