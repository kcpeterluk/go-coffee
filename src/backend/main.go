package main

import (
	"fmt"

	customer "github.com/kcpeterluk/go-coffee/customer/service"
)

func main() {
	customer := customer.NewCustomer("Peter", "Luk")
	fmt.Printf("Hello World %v (%v)\n", customer.Person.FirstName, customer.Person.ID)
}
