/* @dac start
id: how-to-develop-an-activity-definition-in-go
title: How to develop an Activity Definition in Go
label: Activity Definition
description: In the Temporal Go SDK programming model, an Activity Definition is an exportable function or `stuct` method.
*/

/*
In the Temporal Go SDK programming model, an Activity Definition is an exportable function or a `struct` method.
Below is an example of an Activity defined as a Struct method.
*/
package yourapp

import (
	"context"

	"go.temporal.io/sdk/activity"
)

// YourActivityObject is the struct that maintains shared state across Activities.
// If the Worker crashes this Activity object loses its state.
type YourActivityObject struct {
	SharedMessageState *string
	SharedCounterState *int
}

// YourActivityDefinition is your custom Activity Definition.
// An Activity Definiton is an exportable function.
func (a *YourActivityObject) YourActivityDefinition(ctx context.Context, param YourActivityParam) (YourActivityResultObject, error) {
	// Use Acivities for computations or calling external APIs.
	// This is just an example of appending to text and incrementing a counter.
	message := param.ActivityParamX + " World!"
	counter := param.ActivityParamY + 1
	a.SharedMessageState = &message
	a.SharedCounterState = &counter
	result := YourActivityResultObject{
		ResultFieldX: *a.SharedMessageState,
		ResultFieldY: *a.SharedCounterState,
	}
	// Return the results back to the Workflow Execution.
	// The results persist within the Event History of the Workflow Execution.
	return result, nil
}

// PrintSharedState is another custom Activity Definition.
func (a *YourActivityObject) PrintSharedSate(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("The current message is:", *a.SharedMessageState)
	logger.Info("The current counter is:", *a.SharedCounterState)
	return nil
}

/*
An _Activity struct_ can have more than one method, with each method acting as a separate Activity Type.
Activities written as struct methods can use shared struct variables, such as:

- an application level DB pool
- client connection to another service
- reusable utilities
- any other expensive resources that you only want to initialize once per process

Because this is such a common need, the rest of this guide shows Activities written as `struct` methods.
*/

// @dac end
