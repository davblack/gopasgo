package main

import (
	"fmt"
	"os"
	"strconv"
)

// Define a struct for Employee type
type Employee struct {
	Name        string
	GrossIncome float64
}

// Define a method on Employee struct to calculate net income
func (e *Employee) NetIncome() float64 {
	// Prompt user for tax rate
	var taxRate float64
	fmt.Print("Enter tax rate (0.2 for 20%): ")
	fmt.Scanln(&taxRate)

	// Calculate net income based on gross income and tax rate
	return e.GrossIncome * (1 - taxRate)
}

func main() {
	// Parse the command-line argument for gross income
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [gross income]")
		return
	}

	grossIncome, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid input for gross income")
		return
	}

	// Create a new instance of Employee
	emp := Employee{Name: "David", GrossIncome: grossIncome}

	// Call the NetIncome method to calculate net income based on user input
	netIncome := emp.NetIncome()

	// Print out the employees name and net income
	fmt.Printf("Employee %s has a net income of %.2f\n", emp.Name, netIncome)
}
