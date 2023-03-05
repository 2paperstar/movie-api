//go:build !debug

package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	app := setup()
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			log.Println("error", err)
		}
		return err
	})
	fiberLambda = fiberadapter.New(app)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
