package main

import (
	"context"

	"go.temporal.io/sdk/client"
)

func main() {
	verifyState := func(ctx context.Context, scheduleHandle client.ScheduleHandle, paused bool, note string) {
		
	}
}

/* @dacx
id: how-to-pause-a-schedule-in-go
title: How to pause a Schedule in Go
label: Pause Schedule
description: 
lines: 
@dacx */