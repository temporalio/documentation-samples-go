package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
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

	// This schedule ID can be user business logic identifier as well.
	scheduleID := "schedule_" + uuid.New()
	workflowID := "schedule_workflow_" + uuid.New()
	// Create the schedule, start with no spec so the schedule will not run.
	scheduleHandle, err := c.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:   scheduleID,
		Spec: client.ScheduleSpec{},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule.SampleScheduleWorkflow,
			TaskQueue: "schedule",
		},
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
}