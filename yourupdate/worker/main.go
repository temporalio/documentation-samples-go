package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"documentation-samples-go/yourupdate"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()
	w := worker.New(temporalClient, yourupdate.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(yourupdate.YourUpdatableWorkflow)
	w.RegisterWorkflow(yourupdate.UpdatableWorkflowWithValidator)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
