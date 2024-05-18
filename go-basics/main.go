package main

import (
	"fmt"
	"gobasics/imports"
)

func main() {
	tickets := imports.Ticket{
		ID:    "2",
		Event: "Some Event",
	}

	fmt.Printf("%+v", tickets)
	fmt.Printf("helo world")
}
