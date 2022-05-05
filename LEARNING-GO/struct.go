package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Robbie",
		LastName:  "Lawler",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)
}

func whatever() {

}

func Whatever() {

}
