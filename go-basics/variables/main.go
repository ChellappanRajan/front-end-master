package main

import "fmt"

func main() {
	var myName string = "Chellappan"
	myInt := 10
	myFloat := 10.10

	fmt.Printf("Hello everyone my name is %s , my float %f my int %d\n", myName, myFloat, myInt)

	var myOtherFriendName string
	var myOtherInt int
	var myOtherbool bool
	var myOtherFloat float64

	//Zero value by default
	fmt.Printf("hello everyone my friend name %s, my other int %d , my other float %f, my other bool", myOtherFriendName, myOtherInt, myOtherFloat, myOtherbool)
}
