package EmployeeManage

import "fmt"

func main() {
	company := Company{employees: make(map[string]Employee)}

	fte1 := FullTimeEmployee{ID: 1, Name: "Alice", Salary: 300000}
	fte2 := FullTimeEmployee{ID: 2, Name: "Bob", Salary: 350000}

	pte1 := PartTimeEmployee{ID: 3, Name: "Charlie", HourlyRate: 5000, HoursWorked: 20.5}
	pte2 := PartTimeEmployee{ID: 4, Name: "Diana", HourlyRate: 6000, HoursWorked: 15.0}

	company.AddEmployee(fte1)
	company.AddEmployee(fte2)
	company.AddEmployee(pte1)
	company.AddEmployee(pte2)

	fmt.Println("\nEmployee List:")
	company.ListEmployees()
}
