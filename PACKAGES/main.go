package main

// import "github.com/owais34/Building-Modern-web-apps-with-Go/PACKAGES/helpers"

// const numPool = 10

// func CalculateValue(intChan chan int) {
// 	randomNumber = helpers.RandomNumber(numPool)
// 	intChan <- randomNumber
// }
// func main() {
// 	intChan := make(chan int)
// 	defer close(intChan)

// 	go CalculateValue(intChan)

// }

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

// func main() {
// 	myJson := `
// 	[
// 		{
// 			"first_name":"Clark",
// 			"last_name":"Kent",
// 			"hair_color":"black",
// 			"has_dog":true
// 		},
// 		{
// 			"first_name":"Bruce",
// 			"last_name":"Wayne",
// 			"hair_color":"black",
// 			"has_dog":false
// 		}
// 	]
// 	`
// 	var unMarshalled []Person
// 	err := json.Unmarshal([]byte(myJson), &unMarshalled)
// 	if err != nil {
// 		log.Println("error unmarshalling json", err)
// 	}

// 	log.Printf("Unmarshalled : %v", unMarshalled)

// 	//write json from struct
// 	var mySlice []Person
// 	mySlice = append(mySlice, Person{
// 		FirstName: "Owais",
// 		LastName:  "Ahmed",
// 		HairColor: "black",
// 		HasDog:    false,
// 	})
// 	mySlice = append(mySlice, Person{
// 		FirstName: "Bob",
// 		LastName:  "Builder",
// 		HairColor: "brown",
// 		HasDog:    false,
// 	})
// 	newJson, err := json.MarshalIndent(mySlice, "", "   ")
// 	if err != nil {
// 		log.Println("error marshalling")
// 	}

// 	fmt.Print(string(newJson))
// }
