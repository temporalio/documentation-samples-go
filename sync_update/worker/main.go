package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"documentation-samples-go/sync_update"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()
	w := worker.New(temporalClient, sync_update.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(sync_update.YourUpdatableWorkflow)
	w.RegisterWorkflow(sync_update.UpdatableWorkflowWithValidator)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
