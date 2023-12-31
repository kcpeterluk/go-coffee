package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/kcpeterluk/go-coffee/customer/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID        uuid.UUID `bson:"id"`
	FirstName string    `bson:"firstName"`
	LastName  string    `bson:"lastName"`
	Email     string    `bson:"email"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:        c.ID,
		FirstName: c.Person.FirstName,
		LastName:  c.Person.LastName,
		Email:     c.Email,
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetFirstName(m.FirstName)
	c.SetLastName(m.LastName)
	c.SetEmail(m.Email)

	return c
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("coffee")
	customer := db.Collection("customer")

	return &MongoRepository{
		db:       db,
		customer: customer,
	}, nil
}

func (r *MongoRepository) Create(c aggregate.Customer) error {
	_, err := r.customer.InsertOne(context.Background(), NewFromCustomer(c))
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	var c mongoCustomer
	err := r.customer.FindOne(context.Background(), bson.M{"id": id}).Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToAggregate(), nil
}
