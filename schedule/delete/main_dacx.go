package main

import (
	"log"
)

/*
 */

func main() {
	log.Println("Deleting schedule", "ScheduleID", scheduleHandle.GetID())
	err = scheduleHandle.Delete(ctx)
	if err != nil {
		log.Fatalln("Unable to delete schedule", err)
	}
}

/* @dacx
id: delete-schedule-in-go
title: How to Delete a Schedule in Go
label: Delete
description: Use Temporal's Workflow API to delete a Schedule.
lines: 12-15
@dacx */