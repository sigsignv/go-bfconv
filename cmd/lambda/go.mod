module github.com/sigsignv/bfconv/cmd/lambda

go 1.24.0

require (
	github.com/aws/aws-lambda-go v1.52.0
	github.com/sigsignv/bfconv v0.0.0
)

require (
	github.com/mmcdole/gofeed v1.3.0 // indirect
	github.com/mmcdole/goxpp v1.1.1 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

replace github.com/sigsignv/bfconv => ../..
