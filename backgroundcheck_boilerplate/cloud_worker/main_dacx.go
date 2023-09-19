package main

import (
	"log"

	"github.com/joho/godotenv"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"documentation-samples-go/backgroundcheck_boilerplate"
)

/*
To run a Worker Process, define the following steps in code:

- Initialize a Temporal Client.
- Create a new Worker by passing the Client to creation call.
- Register the application's Workflow and Activity functions.
- Call run on the Worker.

In regards to organization, we recommend keeping Worker code separate from Workflow and Activity code.
*/

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Unable to load environment variables from file", err)
	}
	// Initialize a Temporal Client
	// Specify the Namespace in the Client options
	clientOptions := client.Options{
		Namespace: "backgroundcheck_namespace",
	}
	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create a Temporal Client", err)
	}
	defer temporalClient.Close()
	// Create a new Worker
	yourWorker := worker.New(temporalClient, "backgroundcheck-boilerplate-task-queue", worker.Options{})
	// Register Workflows
	yourWorker.RegisterWorkflow(backgroundcheck_boilerplate.BackgroundCheck)
	// Register Acivities
	yourWorker.RegisterActivity(backgroundcheck_boilerplate.SSNTraceActivity)
	// Start the the Worker Process
	err = yourWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start the Worker Process", err)
	}
}

/* @dacx
id: backgroundcheck-boilerplate-cloud-worker
title: Run a dev server Worker
description: Define the code needed to run a Worker Process in Go.
label: Dev server Worker
lines: 1-51
tags: worker, developer guide, temporal client
@dacx */
