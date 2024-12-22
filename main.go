package main

//func main() {
//	library := Library{Books: make(map[string]Book)}
//	reader := bufio.NewReader(os.Stdin)
//	idCounter := 0
//
//	for {
//		fmt.Println("\nMenu:\n1. Add a book\n2. Borrow a book\n3. Return a book\n4. List available books\n5. Exit")
//		fmt.Print("Enter your choice: ")
//		choiceInput, _ := reader.ReadString('\n')
//		choice, err := strconv.Atoi(strings.TrimSpace(choiceInput))
//		if err != nil {
//			fmt.Println("Invalid input, please enter a number between 1 and 5.")
//			continue
//		}
//
//		switch choice {
//		case 1:
//			fmt.Print("Enter book title: ")
//			title, _ := reader.ReadString('\n')
//			fmt.Print("Enter book author: ")
//			author, _ := reader.ReadString('\n')
//
//			book := Book{
//				ID:         strconv.Itoa(idCounter),
//				Title:      strings.TrimSpace(title),
//				Author:     strings.TrimSpace(author),
//				IsBorrowed: false,
//			}
//			library.AddBook(book)
//			idCounter++
//
//		case 2:
//			fmt.Print("Enter book ID to borrow: ")
//			id, _ := reader.ReadString('\n')
//			library.BorrowBook(strings.TrimSpace(id))
//
//		case 3:
//			fmt.Print("Enter book ID to return: ")
//			id, _ := reader.ReadString('\n')
//			library.ReturnBook(strings.TrimSpace(id))
//
//		case 4:
//			library.ListBooks()
//
//		case 5:
//			fmt.Println("Exiting the program. Goodbye!")
//			return
//
//		default:
//			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
//		}
//	}
//}

//func main() {
//	shapes := []Shape{
//		Rectangle{length: 10, width: 5},
//		Circle{radius: 7},
//		Square{length: 4},
//		Triangle{sideA: 3, sideB: 4, sideC: 5},
//	}
//
//	for _, shape := range shapes {
//		PrintShapeDetails(shape)
//	}
//}

//func main() {
//	company := Company{employees: make(map[string]Employee)}
//
//	fte1 := FullTimeEmployee{ID: 1, Name: "Alice", Salary: 300000}
//	fte2 := FullTimeEmployee{ID: 2, Name: "Bob", Salary: 350000}
//
//	pte1 := PartTimeEmployee{ID: 3, Name: "Charlie", HourlyRate: 5000, HoursWorked: 20.5}
//	pte2 := PartTimeEmployee{ID: 4, Name: "Diana", HourlyRate: 6000, HoursWorked: 15.0}
//
//	company.AddEmployee(fte1)
//	company.AddEmployee(fte2)
//	company.AddEmployee(pte1)
//	company.AddEmployee(pte2)
//
//	fmt.Println("\nEmployee List:")
//	company.ListEmployees()
//}

//func main() {
//	account := &BankAccount{
//		AccountNumber: "1234567890",
//		HolderName:    "Adolf",
//		Balance:       0,
//	}
//
//	reader := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Println("\nMenu:\n1. Deposit\n2. Withdraw\n3. Get balance\n4. Exit")
//		fmt.Print("Enter your choice: ")
//		choiceInput, _ := reader.ReadString('\n')
//		choice, err := strconv.Atoi(strings.TrimSpace(choiceInput))
//		if err != nil {
//			fmt.Println("Invalid input, please enter a number between 1 and 4.")
//			continue
//		}
//
//		switch choice {
//		case 1:
//			fmt.Print("Enter deposit amount: ")
//			amountInput, _ := reader.ReadString('\n')
//			amount, err := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
//			if err != nil || amount <= 0 {
//				fmt.Println("Invalid amount. Please enter a positive number.")
//				continue
//			}
//			account.Deposit(amount)
//
//		case 2:
//			fmt.Print("Enter withdrawal amount: ")
//			amountInput, _ := reader.ReadString('\n')
//			amount, err := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
//			if err != nil || amount <= 0 {
//				fmt.Println("Invalid amount. Please enter a positive number.")
//				continue
//			}
//			account.Withdraw(amount)
//
//		case 3:
//			account.GetBalance()
//
//		case 4:
//			fmt.Println("Exiting the program. Goodbye!")
//			return
//
//		default:
//			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
//		}
//	}
//}
