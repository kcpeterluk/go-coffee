package customer

import (
	"github.com/kcpeterluk/go-coffee/customer/aggregate"
)

func NewCustomer(firstName string, lastName string) aggregate.Customer {
	return aggregate.Create(firstName, lastName)
}
