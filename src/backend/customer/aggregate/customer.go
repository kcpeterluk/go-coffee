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

func (c *Customer) SetID(id uuid.UUID) {
	c.ID = id
}

func (c *Customer) SetFirstName(firstName string) {
	if c.Person == nil {
		person := valueobject.NewPerson(firstName, "")
		c.Person = &person
	}
	c.Person.FirstName = firstName
}

func (c *Customer) SetLastName(lastName string) {
	if c.Person == nil {
		person := valueobject.NewPerson("", lastName)
		c.Person = &person
	}
	c.Person.LastName = lastName
}

func (c *Customer) SetEmail(email string) {
	c.Email = email
}
