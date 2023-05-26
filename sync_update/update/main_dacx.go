package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.temporal.io/sdk/client"

	"documentation-samples-go/sync_update"
)

/*
Use the `UpdateWorkflow()` method on an instance of the [Go SDK Temporal Client](https://pkg.go.dev/go.temporal.io/sdk/client#Client) to send an [Update](/concepts/what-is-an-update) to a [Workflow Execution](/workflows#workflow-execution).

The Workflow Id is required, specifying a Run Id is optional.
If only the Workflow Id is supplied (provide an empty string as the Run Id param), the Workflow Execution that is running receives the Signal.
*/

func main() {
	// Exit if an argument is not provided.
	if len(os.Args) != 2 {
		log.Fatalln("Expected a single integer argument")
	}
	// Get the argument from the command line.
	arg := os.Args[1]
	// Ensure the argument is an integer and exit if it is not.
	n, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("The argument must be an integer")
		os.Exit(1)
	}
	// Create a Temporal Client.
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()

	// Set the Update argument values.
	updateArg := sync_update.YourUpdateArg{
		Add: n,
	}
	// Call the UpdateWorkflow API.
	updateHandle, err := temporalClient.UpdateWorkflow(context.Background(), sync_update.YourUpdateWFID, "", sync_update.YourUpdateName, updateArg)
	if err != nil {
		log.Fatalln("Error issuing Update request", err)
		return
	}
	// Get the result of the Update.
	var updateResult sync_update.YourUpdateResult
	err = updateHandle.Get(context.Background(), &updateResult)
	if err != nil {
		log.Fatalln("Update encountered an error", err)
		return
	}
	log.Println("Update succeeded, new total: ", updateResult.Total)
}

/* @dacx
id: how-to-send-an-update-from-a-client-in-go
title: How to send an Update from a Temporal Client in Go
label: Send Update from Client
description: Use the `UpdateWorkflow()` method on an instance of the Go SDK Temporal Client to send an Update to a Workflow Execution.
lines: 15-22, 44-62
@dax */