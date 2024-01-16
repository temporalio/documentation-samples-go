package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"documentation-samples-go/cancel"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()
	// Call the CancelWorkflow API to cancel a Workflow
	// In this call we are relying on the Workflow Id only.
	// But a Run Id can also be supplied to ensure the correct Workflow is Canceled.
	err = temporalClient.CancelWorkflow(context.Background(), cancel.WorkflowId, "")
	if err != nil {
		log.Fatalln("Unable to cancel Workflow Execution", err)
	}
	log.Println("Workflow Execution cancelled", "WorkflowID", cancel.WorkflowId)
}

/* @dacx
id: how-to-request-cancellation-of-workflow-and-activities
title: How to request Cancellation of a Workflow and Activities in Go
label: Request Cancellation
description: Use the Temporal Client's CancelWorkflow API to send a Cancellation Request to the Workflow.
lines: 11, 20-26, 28
@dacx */
