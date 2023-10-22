package valueobject

type Person struct {
	FirstName string
	LastName  string
}

func NewPerson(firstName string, lastName string) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
	}
}