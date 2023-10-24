package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kcpeterluk/go-coffee/customer/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
)

type CustomerRepository interface {
	Create(aggregate.Customer) error
	Get(uuid.UUID) (aggregate.Customer, error)
}
