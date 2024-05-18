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

	for i := 0; i < len(dynamicAnimals); i++ {
		fmt.Println(dynamicAnimals[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range dynamicAnimals {
		fmt.Println(index, value)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
