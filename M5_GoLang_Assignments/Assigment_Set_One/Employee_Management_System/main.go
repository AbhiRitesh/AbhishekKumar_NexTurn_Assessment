package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

const (
	HR = "HR"
	IT = "IT"
	Finance = "Finance"
)

var employees []Employee

func AddEmp(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("Employee ID must be unique")
		}
	}

	if age <= 18 {
		return errors.New("Employee age must be greater than 18")
	}

	employees = append(employees, Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

func searchEmp(searchTerm string) (Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == searchTerm || strings.EqualFold(emp.Name, searchTerm) {
			return emp, nil
		}
	}
	return Employee{}, errors.New("Employee not found")
}

func listEmpByDep(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

func countEmp(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	err := AddEmp(1, "Abhi", 25, HR)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddEmp(2, "Ashish", 30, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddEmp(3, "Deepak", 22, Finance)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddEmp(4, "Sonu", 23, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("All employees added successfully.")

	searchResult, err := searchEmp("Ashish")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Search result:", searchResult)
	}

	fmt.Println("Employees in IT department:", listEmpByDep(IT))

	fmt.Printf("Number of employees in HR: %d\n", countEmp(HR))
}
