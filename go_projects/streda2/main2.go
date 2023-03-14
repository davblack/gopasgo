package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Employee struct {
	income           float64
	name             string
	nonTaxedPart     bool
	underageChildren uint8
	childrenAbove18  uint8
}

func main() {
	questions := []string{"Input your income: ", "Whats your name?", "Do you apply non-taxable item (y/n) ?",
		"How much children under 18 y.o do you have?", "Do you have and live with any children, which is 18-25 years old ?"}
	i := 0
	emp := Employee{}
	for {
		fmt.Println(questions[i])
		input := bufio.NewScanner(os.Stdin) // prepare input scanner input
		input.Scan()                        // wait for input
		m := input.Text()                   // assign input value as string

		switch i {
		case 0:
			if s, err := strconv.ParseFloat(m, 64); err == nil { // check if value is float64 for income
				emp.income = s
				i++
			}
		case 1:
			matching := regexp.MustCompile(`^[a-zA-Z]+|[a-zA-Z]+[a-zA-Z]+$`).MatchString // check if name only consists of letters
			if matching(m) {
				emp.name = m
				i++
			}

		case 2:
			if m == "y" || m == "Y" { // check if values are Y or N or something else
				emp.nonTaxedPart = true
				i++
			} else if m == "N" || m == "n" {
				emp.nonTaxedPart = false
				i++
			}

		case 3:
			if s, err := strconv.ParseUint(m, 10, 8); err == nil { // check if input can be converted to uint8
				emp.underageChildren = uint8(s)
				i++
			}
		case 4:
			if s, err := strconv.ParseUint(m, 10, 8); err == nil {
				emp.childrenAbove18 = uint8(s)
				i++
			}

		default:
			fmt.Printf("something went wrong")
		}

		if i == len(questions) { // end while loop
			break
		}

	}

	//fmt.Println(emp)
	/*
		v := reflect.ValueOf(emp)
		for i := 0; i < v.NumField(); i++ {
			fmt.Println(v.Field(i))
		}
	*/
	emp.netto()                                                                               // call function for receiver
	fmt.Printf("%v your income after tax is: \n%.2f \n", strings.Title(emp.name), emp.income) // print result

}

func (e *Employee) netto() float64 {
	nonTaxAmount := 410.24                                                              // ammount of non taxed income
	afterSocTax := e.income - e.income*0.1340                                           // social + health tax
	TaxRemain := (afterSocTax - nonTaxAmount) * 0.19                                    // taxed remainder
	TaxBonusChildren := float64(e.underageChildren)*140 + float64(e.childrenAbove18)*50 // children bonuus
	e.income = afterSocTax - TaxRemain + TaxBonusChildren                               // final sum
	return e.income
}
