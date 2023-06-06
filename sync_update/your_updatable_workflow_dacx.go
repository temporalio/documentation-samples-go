package sync_update

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

/*
In Go, you define an Update type, also known as an Update name, as a `string` value.
You must ensure the arguments and result are [serializable](/concepts/what-is-a-data-converter).
When sending and receiving the Update, use the Update name as an identifier.
The name does not link to the data type(s) sent with the Update.
Ensure that every Workflow listening to the same Update name can handle the same Update arguments.
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
Register an Update handler for a given name using the [SetUpdateHandler](https://pkg.go.dev/go.temporal.io/sdk/workflow#SetUpdateHandler) API from the `go.temporal.io/sdk/workflow` package.
The handler function can accept multiple serializable input parameters, but we recommend using only a single parameter.
This practice allows you to add fields in future versions while maintaining backward compatibility.
You can optionally include a `workflow.Context` parameter in the first position of the function.
The function can return either a serializable value with an error or just an error.

Update handlers, unlike Query handlers, can observe and mutate Workflow state safely.
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
Validate certain aspects of the data sent to the Workflow using an Update validator function.
For instance, a counter Workflow might never want to accept a non-positive number.
Invoke the `SetUpdateHandlerWithOptions` API and define a validator function as one of the options.

When you use a Validator function, the Worker receives the Update first, before any Events are written to the Event History.
If the Update is rejected, it's not recorded in the Event History.
If it's accepted, the `WorkflowExecutionUpdateAccepted` Event occurs.
Afterwards, the Worker executes the accepted Update and, upon completion, a `WorkflowExecutionUpdateCompleted` Event gets written into the Event History.
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
label: Update type
description: An Update type, also called an Update name, is a string value.
lines: 10-19, 67,69,75,82
@dacx */

/* @dacx
id: how-to-handle-an-update-in-go
title: How to handle an Update in Go
label: Handle Update
description: Use the SetUpateHandler API from the go.temporal.io/sdk/workflow package to register an Update Handler for a given name.
lines: 53-75, 82
@dacx */

/* @dacx
id: how-to-set-an-update-validator-function-in-go
title: How to set an Update validator function in go
label: Validator function
description: Use the SetUpdateHandlerWithOptions API and pass it a validator function to validate inputs.
lines: 84-103, 109-114, 121, 123-133
@dacx */
