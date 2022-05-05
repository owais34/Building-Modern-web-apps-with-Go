package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])

	myMap2 := make(map[string]User)
	myMap2["son"] = User{FirstName: "Joe", LastName: "Rogan"}

}
