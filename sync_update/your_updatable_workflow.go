package sync_update

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

/*
In Go, an Update type, also called an Update name, is a `string` value.
The arguments and result must be [serializable](/concepts/what-is-a-data-converter).
*/

const YourUpdateName = "your_update_name"

const YourValidatedUpdateName = "your_validated_update_name"

const TaskQueueName = "your_updatable_workflow"

const YourUpdateWFID = "updatable_workflow"

const YourValidUpdateWFID = "validating_updatable_workflow"

type YourUpdateArg struct {
	Add int
}

type YourUpdateResult struct {
	Total int
}

type WFParam struct {
	StartCount int
}

type WFResult struct {
	EndTotal int
}

/*
Use the [SetUpdateHandler](https://pkg.go.dev/go.temporal.io/sdk/workflow#SetUpdateHandler) API from the `go.temporal.io/sdk/workflow` package to register an Update handler for a given name.
*/

func YourUpdatableWorkflow(ctx workflow.Context, param WFParam) (WFResult, error) {
	counter := param.StartCount
	workflow.SetUpdateHandler(ctx, YourUpdateName, func(arg YourUpdateArg) (YourUpdateResult, error) {
		counter += arg.Add
		result := YourUpdateResult{
			Total: counter,
		}
		return result, nil
	})
	// Sleep for 60 seconds to have time to send Updates
	workflow.Sleep(ctx, 60*time.Second)
	endTotal := WFResult{
		EndTotal: counter,
	}
	return endTotal, nil
}

/*
In the preceding example, the Workflow code uses `workflow.SetUpdateHandler` to register a function to handle Workflow Updates.
The function can take multiple serializable input parameters, although we recommend that you use only a single parameter to allow for fields to be added in future versions while retaining backward compatibility.
The function can optionally take a `workflow.Context` parameter in the first position.
The function returns either a serializable value and an error or just an error.

Unlike Query handlers, Update handlers can safely observe and mutate Workflow state.
*/

func UpdatableWorkflowWithValidator(ctx workflow.Context, param WFParam) (WFResult, error) {
	counter := param.StartCount
	workflow.SetUpdateHandlerWithOptions(
		ctx, YourValidatedUpdateName,
		func(arg YourUpdateArg) (YourUpdateResult, error) {
			counter += arg.Add
			result := YourUpdateResult{
				Total: counter,
			}
			return result, nil
		},
		workflow.UpdateHandlerOptions{Validator: isPositive},
	)
	// Sleep for 60 seconds to have time to send Updates
	workflow.Sleep(ctx, 60*time.Second)
	endTotal := WFResult{
		EndTotal: counter,
	}
	return endTotal, nil
}

func isPositive(ctx workflow.Context, i int) error {
	log := workflow.GetLogger(ctx)
	if i < 0 {
		log.Debug("Rejecting negative number, positive integers only", "update value:", i)
		return fmt.Errorf("addend must be a positive integer (%v)", i)
	}
	log.Debug("Accepting update", "update value:", i)
	return nil
}

/* @dacx
id: how-to-handle-an-update-in-go
title: How to handle an Update in Go
sidebar_label: Handle Update
description: Use the `SetUpateHandler` API from the `go.temporal.io/sdk/workflow` package to register an Update Handler for a given name.
tags:
  - go
  - how-to
@dacx */

/* @dacx
id: how-to-define-an-update-type-in-go
title: How to define an Update Type in Go
sidebar_label: Update type
description: An Update type, also called an Update name, is a `string` value.
tags:
  - go
  - how-to
@dacx */
