package main

import (
	"log"

	"go.temporal.io/sdk/client"
)

/*
 */

func main() {
	// Manually trigger the schedule once
	log.Println("Manually triggering schedule", "ScheduleID", scheduleHandle.GetID())

	err = scheduleHandle.Trigger(ctx, client.ScheduleTriggerOptions{
		Overlap: enums.SCHEDULE_OVERLAP_POLICY_ALLOW_ALL,
	})
	if err != nil {
		log.Fatalln("Unable to trigger schedule", err)
	}
}

/* @dacx
id: describe-schedule-in-go
title: How to Describe a Schedule in Go
label: Describe
description: Use Temporal's Workflow API to display information about existing Schedules.
lines: 16-18
@dacx */