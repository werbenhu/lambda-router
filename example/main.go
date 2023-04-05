package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	router "github.com/werbenhu/lambda-router"
)

func main() {

	router.Get("/", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("match path /\n")
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	router.Get("/ttt/", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("match path /ttt\n")
		fmt.Printf("params: %+v\n", request.QueryStringParameters)
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	router.Get("/ttt/:abc", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("match path /ttt rest api with value \n")
		fmt.Printf("params: %+v\n", request.QueryStringParameters)
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	group := router.NewGroup("test")
	group.Get("/ttt/:abc", func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Printf("match path /test/ttt rest api with value \n")
		fmt.Printf("params: %+v\n", request.QueryStringParameters)
		return events.APIGatewayProxyResponse{Body: "GET WERBEN IN ROUTER", StatusCode: 200}, nil
	})

	ctx := context.Background()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/",
	}
	ret, err := router.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/ttt"
	ret, err = router.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/ttt/werben"
	ret, err = router.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/test/ttt/123"
	ret, err = router.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)

	request.Path = "/t"
	ret, err = router.Handler(ctx, request)
	fmt.Printf("ret:%+v, err:%s\n", ret, err)
}
