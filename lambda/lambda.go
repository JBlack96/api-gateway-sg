package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, request APIRequest) (APIResponse, error) {
	b := RequestBody{}
	json.Unmarshal([]byte(request.Body), &b)

	return Response(200, fmt.Sprintf("Hello %s!", b.Name)), nil
}

func main() {
	lambda.Start(HandleRequest)
}

// APIRequest type alias for readability
type APIRequest = events.APIGatewayProxyRequest

// APIResponse type alias for readability
type APIResponse = events.APIGatewayProxyResponse

// Response func
func Response(status int, body string) APIResponse {
	resp := APIResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}

	if body != "" {
		resp.Body = body
	}

	return resp
}
