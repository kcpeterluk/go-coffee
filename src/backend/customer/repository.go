package customer

import (
	"errors"

	"github.com/kcpeterluk/go-coffee/customer/aggregate"
)

var (
	ErrFailedToAddCustomer = errors.New("failed to add customer")
)

type CustomerRepository interface {
	Create(aggregate.Customer) error
}
