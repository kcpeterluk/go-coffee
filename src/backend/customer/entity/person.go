package entity

import (
	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}

func NewPerson(firstName string, lastName string) Person {
	return Person{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
	}
}