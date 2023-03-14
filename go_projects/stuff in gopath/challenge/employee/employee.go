package employee

import "fmt"

// Account represents an account for the online store
type Account struct {
	FirstName string
	LastName  string
}

// ChangeName changes the name of the account owner
func (a *Account) ChangeName(firstName, lastName string) {
	a.FirstName = firstName
	a.LastName = lastName
}

// String returns the account owner's name in "first name last name" format
func (a *Account) String() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

// Employee represents an employee of the online store
type Employee struct {
	*Account
	Credits float64
}

// NewEmployee creates a new Employee object with the specified name and credits
func NewEmployee(firstName, lastName string, credits float64) *Employee {
	return &Employee{
		Account: &Account{FirstName: firstName, LastName: lastName},
		Credits: credits,
	}
}

// AddCredits adds the specified amount of credits to the employee's account
func (e *Employee) AddCredits(amount float64) {
	e.Credits += amount
}

// RemoveCredits removes the specified amount of credits from the employee's account
func (e *Employee) RemoveCredits(amount float64) {
	if e.Credits < amount {
		fmt.Println("Insufficient credits")
		return
	}
	e.Credits -= amount
}

// CheckCredits returns the employee's current credit balance
func (e *Employee) CheckCredits() float64 {
	return e.Credits
}
