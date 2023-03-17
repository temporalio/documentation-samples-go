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
	workflowOptions := client.StartWorkflowOptions{
		ID:        cancel.WorkflowId,
		TaskQueue: cancel.TaskQueueName,
	}
	we, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, cancel.YourWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
