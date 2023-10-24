package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kcpeterluk/go-coffee/customer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoUri := os.Getenv("MONGODB_URI")

	cs, err := customer.NewCustomerService(customer.WithMongoCustomerRepository(mongoUri))
	if err != nil {
		panic(err)
	}

	id, err := cs.Create("Peter", "Luk", "peter.luk@example.com")
	if err != nil {
		panic(err)
	}

	c, err := cs.Get(id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %v\n", c.Person.FirstName)
}
