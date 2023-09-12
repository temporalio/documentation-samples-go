package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"documentation-samples-go/backgroundcheckboilerplate"
)

/*
We recommend keeping Worker code separate from Workflow and Activity code.
*/

func main() {
	// Initialize a Temporal Client
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create a Temporal Client", err)
	}
	defer temporalClient.Close()
	// Create a new Worker
	yourWorker := worker.New(temporalClient, "background-check-boilerplate-task-queue", worker.Options{})
	// Register Workflows
	yourWorker.RegisterWorkflow(backgroundcheckboilerplate.BackgroundCheck)
	// Register Acivities
	yourWorker.RegisterActivity(backgroundcheckboilerplate.SSNTraceActivity)
	// Start the the Worker Process
	err = yourWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start the Worker Process", err)
	}
}
