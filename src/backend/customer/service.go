package customer

import (
	"context"

	"github.com/kcpeterluk/go-coffee/customer/aggregate"
	mongo "github.com/kcpeterluk/go-coffee/customer/persistence"
)

type CustomerConfiguration func(cs *CustomerService) error

type CustomerService struct {
	customers CustomerRepository
}

func NewCustomerService(cfgs ...CustomerConfiguration) (*CustomerService, error) {
	cs := &CustomerService{}

	for _, cfg := range cfgs {
		if err := cfg(cs); err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func WithMongoCustomerRepository(connectionString string) CustomerConfiguration {
	return func(cs *CustomerService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}

		cs.customers = cr
		return nil
	}
}

func (cs *CustomerService) Create(firstName string, lastName string, email string) error {
	c := aggregate.NewCustomer(firstName, lastName, email)
	if err := cs.customers.Create(c); err != nil {
		return err
	}

	return nil
}