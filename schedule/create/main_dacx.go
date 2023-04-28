package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
)

/* @dacx
Schedules are created with the `create` action.
For each new Schedule, tbe Temporal Server generates a unique Schedule ID.

To create a Schedule in Go, use `ScheduleClient().Create()` on the [Client](/concepts/what-is-the-temporal-client).
Schedules must be initialized with a Schedule ID, [Spec](/concepts/what-is-a-schedule#spec), and [Action](/concepts/what-is-a-schedule#action) to perform.
Enter these values in `client.Schedule.Options{}`.
 @dacx */

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
/* @dacx
id: create-schedule-in-go
title: How to Create a Schedule in Go
label: Create
description: Use Temporal's Workflow API to create a Schedule.
lines: 11-15, 32-35, 16-17, 35-42
@dacx */