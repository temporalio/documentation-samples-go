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

/* @dacx
id: backfill-schedule-in-go
title: How to Backfill a Schedule in Go
label: Backfill
description: Use Temporal's Workflow API to execute Schedules ahead of time.
lines: 
@dacx */