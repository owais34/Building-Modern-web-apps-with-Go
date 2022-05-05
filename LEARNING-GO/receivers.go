package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Jogn"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println("myVar set to ", myVar.printFirstName())
	log.Println("myVar2 set to ", myVar2.printFirstName())
}
