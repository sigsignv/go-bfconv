package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	return `{ "kind": "Hello, World!" }`, nil
}

func main() {
	lambda.Start(hello)
}
