package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/client/schedule"
)

func main() {
    ctx := context.Background()
    temporalClient, err := client.Dial(client.Options{
        HostPort: client.DefaultHostPort,
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal Client", err)
    }
    defer temporalClient.Close()

    // Create Schedule and Workflow IDs
	scheduleID := "schedule_" + uuid.New()
	workflowID := "schedule_workflow_" + uuid.New()
	// Create the schedule.
	scheduleHandle, err := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:   scheduleID,
		Spec: client.ScheduleSpec{},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule.ScheduleWorkflow,
			TaskQueue: "schedule",
		},
	})
	if err != nil {
		log.Fatalln("Unable to create schedule", err)
	}
    log.Println("Schedule created", "ScheduleID", scheduleID)
	_, _ = scheduleHandle.Describe(ctx)
}

/*
Schedules are initiated with the `create` call.
The user generates a unique Schedule ID for each new Schedule.

To create a Schedule in Go, use `Create()` on the [Client](/concepts/what-is-the-temporal-client).
Schedules must be initialized with a Schedule ID, [Spec](/concepts/what-is-a-schedule#spec), and [Action](/concepts/what-is-a-schedule#action) in `client.ScheduleOptions{}`.
*/

/* @dacx
id: how-to-create-a-schedule-in-go
title: How to create a Schedule in Go
label: Create Schedule
description: Create a Schedule for a Workflow in Go.
lines: 12, 22-34, 41
@dacx */