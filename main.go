package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func handlerLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("Hello, %s. You are %d years old.", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(handlerLambdaEvent)
}
