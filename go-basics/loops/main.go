package main

import (
	"fmt"
	"slices"
)

func main() {
	//Array fixed in size
	animals := [2]string{}

	animals[0] = "Dog"
	animals[1] = "Cat"

	fmt.Println(animals)

	//slice dynamic in size
	dynamicAnimals := []string{"Dog", "Cat"}

	dynamicAnimals = append(dynamicAnimals, "Tiger")

	dynamicAnimals = slices.Delete(dynamicAnimals, 1, 2)

	fmt.Println(dynamicAnimals)

}
