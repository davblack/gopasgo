package main

import (
	"challenge/employee"
	"fmt"
)

func main() {
	// Create a new employee
	employee := employee.NewEmployee("John", "Doe", 100.0)

	// Print the employee's name
	fmt.Println("Employee name:", employee.Account)

	// Change the employee's name
	employee.Account.ChangeName("Jane", "Doe")

	// Print the employee's name again
	fmt.Println("Employee name:", employee.Account)

	// Add credits to the employee's account
	employee.AddCredits(50.0)

	// Check the employee's balance
	balance := employee.CheckCredits()
	fmt.Println("Employee balance:", balance)

	// Remove credits from the employee's account
	employee.RemoveCredits(25.0)

	// Check the employee's balance again
	balance = employee.CheckCredits()
	fmt.Println("Employee balance:", balance)
}
