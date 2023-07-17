package main

import (
	"context"
	"log"
	"time"

	"github.com/temporalio/documentation-samples-go/schedule"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a Workflow to backfill
	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

	workflowID := "schedule_workflow_"
	// create paused Workflow
	now := time.Now()
	scheduleHandle, _ := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: "backfill-schedule",
		Spec: client.ScheduleSpec{
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: time.Minute,
				},
			},
		},
		Action: &client.ScheduleWorkflowAction{
			ID:        workflowID,
			Workflow:  schedule.ScheduleWorkflow,
			TaskQueue: "schedule",
		},
		Paused: true,
	})

	err = scheduleHandle.Backfill(ctx, client.ScheduleBackfillOptions{
		Backfill: []client.ScheduleBackfill{
			{
				Start:   now.Add(-4 * time.Minute),
				End:     now.Add(-2 * time.Minute),
				Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
			},
			{
				Start:   now.Add(-2 * time.Minute),
				End:     now,
				Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
			},
		},
	})
	if err != nil {
		log.Fatalln("Unable to Backfill Schedule", err)
	}

}

/*
Backfilling a Schedule executes [Workflow Tasks](/concepts/what-is-a-workflow-task) ahead of the Schedule's specified time range.
This is useful for executing a missed or delayed Action, or for testing the Workflow ahead of time.

To backfill a Schedule in Go, use `Backfill()` on `ScheduleHandle`.
Specify the start and end times to execute the Workflow, along with the overlap policy.
*/

/* @dacx
id: how-to-backfill-a-schedule-in-go
title: How to backfill a Schedule in Go
label: Backfill Schedule
description: Backfill a Schedule to execute a Workflow at a later time.
lines: 14, 48-61, 63, 64-70
@dacx */
