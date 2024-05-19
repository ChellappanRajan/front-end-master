package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func handleRequest(event MyEvent) (string, error) {

	if event.Username == "" {
		return "", fmt.Errorf("username can not be empty")
	}
	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	lambda.Start(handleRequest)
}
