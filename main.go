package main

import (
	"context"
	"fmt"

	"github.com/yashjawale/go-microservice/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}