package aggregate

import (
	"github.com/kcpeterluk/go-coffee/customer/entity"
)

type Customer struct {
	Person *entity.Person
}

func Create(firstName string, lastName string) Customer {
	person := entity.NewPerson(firstName, lastName)

	return Customer{
		Person: &person,
	}
}
