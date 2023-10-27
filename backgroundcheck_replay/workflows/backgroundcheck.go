package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"documentation-samples-go/backgroundcheck_replay/activities"
)

// BackgroundCheck is your custom Workflow Definition.
func BackgroundCheck(ctx workflow.Context, param string) (string, error) {
	// Sleep for 1 minute
	workflow.GetLogger(ctx).Info("Sleeping for 1 minute...")
	// highlight-next-line
	err := workflow.Sleep(ctx, 1*time.Minute)
	if err != nil {
		return err
	}
	workflow.GetLogger(ctx).Info("Finished sleeping")
	// Define the Activity Execution options
	// StartToCloseTimeout or ScheduleToCloseTimeout must be set
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	// Execute the Activity synchronously (wait for the result before proceeding)
	var ssnTraceResult string
	// highlight-next-line
	err := workflow.ExecuteActivity(ctx, activities.SSNTraceActivity, param).Get(ctx, &ssnTraceResult)
	if err != nil {
		return "", err
	}
	// Make the results of the Workflow available
	return ssnTraceResult, nil
}

/* @dacx
id: add-sleep-for-one-minute
title: Add a call to sleep
description: Add a call to sleep for one minute to the beginning of the Workflow.
label: Add sleep call
lines: 71-87
tags:
- testing
- replay test
- replayer
@dacx */
