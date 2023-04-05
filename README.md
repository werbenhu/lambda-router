# Router
aws apigateway lambda router for golang


# Example:
```
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	router "github.com/werbenhu/lambda-router"
)

func init() {
	router.Get("/test", test)
	router.Get("/test/:name", testWithName)

	group := router.NewGroup("aaa")
	group.Get("/bbb", groupWithName)
}

func test(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Headers:%+v\n", request.Headers)
	fmt.Printf("params:%+v\n", request.QueryStringParameters)
	return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
}

func testWithName(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Headers:%+v\n", request.Headers)
	fmt.Printf("params:%+v\n", request.QueryStringParameters)
	return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
}

func groupWithName(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Headers:%+v\n", request.Headers)
	fmt.Printf("params:%+v\n", request.QueryStringParameters)
	return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
}

func main() {
	lambda.Start(router.Handler)
}


```
