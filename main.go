package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginlambda.GinLambda

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		r := gin.Default()
		r.GET("/test", test)
		ginLambda = ginlambda.New(r)
	}
	return ginLambda.Proxy(request)
}

func test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "test!!"})
}
