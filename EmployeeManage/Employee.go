package EmployeeManage

import (
	"fmt"
	"strconv"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

type Company struct {
	employees map[string]Employee
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-Time EmployeeManage: ID=%d, Name=%s, Salary=%d Tenge", fte.ID, fte.Name, fte.Salary)
}

func (pte PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-Time EmployeeManage: ID=%d, Name=%s, Hourly Rate=%d Tenge, Hours Worked=%.2f", pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked)
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.employees[strconv.FormatUint(e.ID, 10)] = emp
	case PartTimeEmployee:
		c.employees[strconv.FormatUint(e.ID, 10)] = emp
	default:
		fmt.Println("Unknown employee type")
	}
	fmt.Println("EmployeeManage added successfully")
}

func (c *Company) ListEmployees() {
	if len(c.employees) == 0 {
		fmt.Println("No employees found")
		return
	}
	for id, emp := range c.employees {
		fmt.Printf("ID: %s, %s\n", id, emp.GetDetails())
	}
}
