package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	log.Printf("Body size = %d.\n", len(request.Body))

	log.Println("Headers:")
	for key, value := range request.Headers {
		log.Printf("    %s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{Body:"Hello World", StatusCode:http.StatusOK}, nil
}
