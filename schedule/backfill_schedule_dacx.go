package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

/*
 */
 func main() {
	ctx := context.Background()
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
 }