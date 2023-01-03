// Created with Strapit
package main

import (
	"context"
	"log"
	"os"

	"github.com/googlecloudplatform/ezcx"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	ctx := context.Background()
	lg := log.Default()
	server := ezcx.NewServer(ctx, ":"+PORT, lg)
	server.HandleCx("/hello", CxHelloWorld)

	server.ListenAndServe(ctx)
}

func CxHelloWorld(res *ezcx.WebhookResponse, req *ezcx.WebhookRequest) error {
	res.AddTextResponse("Hello, Dialogflow CX World!")
	return nil
}
