package main

import (
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "owl", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}

	animalNames := make(map[string]string)
	animalNames["dog"] = "fluffy"
	animalNames["cat"] = "missy"
	animalNames["horse"] = "sherlie"

	for animal, name := range animalNames {
		log.Println(animal, name)
	}
}
