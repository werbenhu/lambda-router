# Router
aws apigateway lambda router for golang


# Example:
```
package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/werbenhu/router"
)

var r *router.Router

func init() {
	r = router.New()
	r.Get("/test", test)
	r.Get("/test/:name", testWithName)
}

func testWithName(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters[":name"]
	resp := map[string]interface{}{
		"name":   name,
		"path":   request.Path,
		"method": request.HTTPMethod,
		"params": request.QueryStringParameters,
	}

	body, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func test(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := map[string]interface{}{
		"path":   request.Path,
		"method": request.HTTPMethod,
		"params": request.QueryStringParameters,
	}

	body, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(r.Handler)
}

```