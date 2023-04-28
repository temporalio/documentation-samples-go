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
// Update the schedule with a spec so it will run periodically,
log.Println("Updating schedule", "ScheduleID", scheduleHandle.GetID())
err = scheduleHandle.Update(ctx, client.ScheduleUpdateOptions{
	DoUpdate: func(schedule client.ScheduleUpdateInput) (*client.ScheduleUpdate, error) {
		schedule.Description.Schedule.Spec = &client.ScheduleSpec{
			// Run the schedule at 5pm on Friday
			Calendars: []client.ScheduleCalendarSpec{
				{
					Hour: []client.ScheduleRange{
						{
							Start: 17,
						},
					},
					DayOfWeek: []client.ScheduleRange{
						{
							Start: 5,
						},
					},
				},
			},
			// Run the schedule every 5s
			Intervals: []client.ScheduleIntervalSpec{
				{
					Every: 5 * time.Second,
				},
			}
		}
	}
}
}

/* @dacx
id: describe-schedule-in-go
title: How to Describe a Schedule in Go
label: Describe
description: Use Temporal's Workflow API to display information about existing Schedules.
lines: 17-19
@dacx */