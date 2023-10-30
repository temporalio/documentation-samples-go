// dacx
// CAUTION! Do not use this code!
package workflows

import (
	"math/rand"
	"time"

	"go.temporal.io/sdk/workflow"

	"documentation-samples-go/backgroundcheck_replay/activities"
)

/*
Referred to as "intrinsic non-determinism" this kind of "bad Workflow code logic" can prevent the Workflow code from executing to completion because the Workflow can take a different code path than the one expected from the Event History.

The following are some common operations that **can't** be done inside of a Wokflow Definition:

- Generating and relying on random numbers (Use Activites instead).
- Accessing / mutating external systems or state (use Activities instead).
- Relying on system time (use Workflow.now() instead).
- Working directly with threads or goroutines (use Workflow.go() instead).
- Iterating over data structures with unknown ordering.
- Storing or evaluating the run Id.

One way to produce a non-deterministic error is to sleep for a random amount of time inside the Workflow.
*/

// BackgroundCheckNonDeterministic is an anti-pattern Workflow Definition
func BackgroundCheckNonDeterministic(ctx workflow.Context, param string) (string, error) {
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	var ssnTraceResult string
	// highlight-start
	// CAUTION, the following code is an anti-pattern showing what NOT to do
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// highlight-end
	err := workflow.ExecuteActivity(ctx, activities.SSNTraceActivity, param).Get(ctx, &ssnTraceResult)
	if err != nil {
		return "", err
	}
	return ssnTraceResult, nil
}

/* @dacx
id: intrinsic-non-determinism
title: Intrinsic non-deterministic logic
description: This kind of logic prevents the Workflow code from executing to completion because the Workflow can take a different code path than the one expected from the Event History.
label: intrinsic-non-deterministic-logic
lines: 3-45
tags:
- tests
- replay
- event history
@dacx */
