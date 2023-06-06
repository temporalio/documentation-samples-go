package sync_update

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

/*
In Go, an Update type, also called an Update name, is a `string` value.
The arguments and result must be [serializable](/concepts/what-is-a-data-converter).
The Update name is used to as an identifier when the Update is sent and received.
The name is not tied to the data type(s) that the are sent with the Update.
Make sure that each Workflow listening to the same Update name can handle the same Update arguments.
*/

// YourUpdateName holds a string value used to correlate Updates.
const YourUpdateName = "your_update_name"

// YourValidatedUpdateName is the name of an Update.
const YourValidatedUpdateName = "your_validated_update_name"

// TaskQueueName is the name of the Task Queue.
const TaskQueueName = "your_updatable_workflow"

// YourUpdateWFID is the Id used for the YourUpdatableWorkflow execution.
const YourUpdateWFID = "updatable_workflow"

// YourValidUpdateWFID is the Id used for the UpdatableWorkflowWithValidator execution.
const YourValidUpdateWFID = "validating_updatable_workflow"

// YourUpdateArg defines the structure of the Update argument.
type YourUpdateArg struct {
	Add int
}

// YourUpdateResult defines the structure of the Update result.
type YourUpdateResult struct {
	Total int
}

// WFParam defines the structure of thw Workflow argument.
type WFParam struct {
	StartCount int
}

// WFResult defines the structure of the Worfklow result.
type WFResult struct {
	EndTotal int
}

/*
Use the [SetUpdateHandler](https://pkg.go.dev/go.temporal.io/sdk/workflow#SetUpdateHandler) API from the `go.temporal.io/sdk/workflow` package to register an Update handler for a given name.
The handler function can take multiple serializable input parameters, although we recommend that you use only a single parameter to allow for fields to be added in future versions while retaining backward compatibility.
The function can optionally take a `workflow.Context` parameter in the first position.
The function can return either a serializable value and an error or just an error.

Unlike Query handlers, Update handlers can safely observe and mutate Workflow state.
*/

// YourUpdatableWorkflow is a Workflow Definition.
// This Workflow sets an Update handler and then sleeps for a minute.
// After setting the Update hanlder it sleeps for 1 minutue.
// Updates can be sent to the Workflow during this time.
func YourUpdatableWorkflow(ctx workflow.Context, param WFParam) (WFResult, error) {
	counter := param.StartCount
	workflow.SetUpdateHandler(ctx, YourUpdateName, func(arg YourUpdateArg) (YourUpdateResult, error) {
		counter += arg.Add
		result := YourUpdateResult{
			Total: counter,
		}
		return result, nil
	})
	// Sleep for 60 seconds to have time to send Updates.
	workflow.Sleep(ctx, 60*time.Second)
	endTotal := WFResult{
		EndTotal: counter,
	}
	return endTotal, nil
}

/*
Use an Update validator function to validate certain aspects of the data sent to the Workflow.
For example, a Workflow that acts as a counter may never want to be fed a non-positive number.
Use the `SetUpdateHanlderWithOptions` API and specify a validator function as one of the options.

When a Validator function is used, the Update is sent to the Worker first before any Events are written to the Event History.
A rejected Update is not written to the Event History.
An accepted Update results in the `WorkflowExecutionUpdateAccepted` Event.
The accepted Update is then executed on the Worker and, upon completion, causes a `WorkflowExecutionUpdateCompleted` Event to be written to the Event History.
*/

// UpdatableWorkflowWithValidator is a Workflow Definition.
// This Workflow Definition has an Update handler that uses the isPositive() validator function.
// After setting the Update hanlder it sleeps for 1 minutue.
// Updates can be sent to the Workflow during this time.
func UpdatableWorkflowWithValidator(ctx workflow.Context, param WFParam) (WFResult, error) {
	counter := param.StartCount
	if err := workflow.SetUpdateHandlerWithOptions(
		ctx, YourValidatedUpdateName,
		func(arg YourUpdateArg) (YourUpdateResult, error) {
			counter += arg.Add
			result := YourUpdateResult{
				Total: counter,
			}
			return result, nil
		},
		// Set the isPositive validator.
		workflow.UpdateHandlerOptions{Validator: isPositive},
	); err != nil {
		return WFResult{}, err
	}
	// Sleep for 60 seconds to have time to send Updates.
	workflow.Sleep(ctx, 60*time.Second)
	endTotal := WFResult{
		EndTotal: counter,
	}
	return endTotal, nil
}

// isPositive is a validator function.
// It returns an error if the int value is below 1.
func isPositive(ctx workflow.Context, u YourUpdateArg) error {
	log := workflow.GetLogger(ctx)
	if u.Add < 1 {
		log.Debug("Rejecting non-positive number, positive integers only", "update value:", u.Add)
		return fmt.Errorf("addend must be a positive integer (%v)", u.Add)
	}
	log.Debug("Accepting update", "update value:", u.Add)
	return nil
}

/* @dacx
id: how-to-define-an-update-type-in-go
title: How to define an Update Type in Go
sidebar_label: Update type
description: An Update type, also called an Update name, is a string value.
lines: 10-19, 66,68,74,81
@dacx */

/* @dacx
id: how-to-handle-an-update-in-go
title: How to handle an Update in Go
sidebar_label: Handle Update
description: Use the SetUpateHandler API from the go.temporal.io/sdk/workflow package to register an Update Handler for a given name.
lines: 53-74, 81
@dacx */

/* @dacx
id: how-to-set-an-update-validator-function-in-go
title: How to set an Update validator function in go
label: Validator function
description: Use the SetUpdateHandlerWithOptions API and pass it a validator function to validate inputs.
lines: 83-102, 108-113, 120, 122-132
@dacx */
