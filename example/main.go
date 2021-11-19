package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/werbenhu/router"
)

func main() {
	r := router.New()

	r.Get("/", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("/\n")
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	r.Get("/ttt/:abc", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("/ttt\n")
		fmt.Printf("%+v\n", request.QueryStringParameters)
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	ctx := context.Background()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/",
	}
	ret, err := r.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/ttt/werben"
	ret, err = r.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/t"
	ret, err = r.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)
}
