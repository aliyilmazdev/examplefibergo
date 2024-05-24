//go:build netlify
// +build netlify

package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/adapter/lambda"
	"github.com/gofiber/fiber/v2"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    app := fiber.New()

    app.Get("/api", func(c *fiber.Ctx) error {
        return c.SendString("Merhaba, Fiber!")
    })

    return lambda.FiberHandler(app)(request)
}

func main() {
    lambda.Start(handler)
}
