package aggregate

import (
	"github.com/google/uuid"
	"github.com/kcpeterluk/go-coffee/customer/valueobject"
)

type Customer struct {
	ID     uuid.UUID
	Person *valueobject.Person
	Email  string
}

func NewCustomer(firstName string, lastName string, email string) Customer {
	person := valueobject.NewPerson(firstName, lastName)

	return Customer{
		ID:     uuid.New(),
		Person: &person,
		Email:  email,
	}
}
