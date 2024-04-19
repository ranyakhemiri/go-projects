package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"how old are you?`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaEvent(Event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", Event.Name, Event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
