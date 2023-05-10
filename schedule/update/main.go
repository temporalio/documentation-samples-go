package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	temporalClient, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client", err)
	}
	defer temporalClient.Close()

	scheduleHandle, _ := temporalClient.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID: "update-schedule",
		Spec: client.ScheduleSpec{},
		Action: &client.ScheduleWorkflowAction{},
		Paused: true,
	})

	updateSchedule := func(input client.ScheduleUpdateInput) (*client.ScheduleUpdate, error) {
		return &client.ScheduleUpdate{
			Schedule:  &input.Description.Schedule,
		}, nil
	}

	_ = scheduleHandle.Update(ctx, client.ScheduleUpdateOptions{
		DoUpdate: updateSchedule,
	})
}

/*
Updating a Schedule changes the configuration of an existing Schedule.

*/

/* @dacx
id: how-to-update-a-schedule-in-go
title: How to update a Schedule in Go
label: Update Schedule
description: Update the configuration of a Schedule.
lines: 
@dacx */