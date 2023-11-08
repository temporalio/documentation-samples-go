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
Referred to as "intrinsic non-determinism" this kind of "bad" Workflow code can prevent the Workflow code from completing because the Workflow can take a different code path than the one expected from the Event History.

The following are some common operations that **can't** be done inside of a Workflow Definition:

- Generate and rely on random numbers (Use Activites instead).
- Accessing / mutating external systems or state.
  This includes calling an external API, conducting a file I/O operation, talking to another service, etc. (use Activities instead).
- Relying on system time.
  - Use `workflow.Now()` as a replacement for `time.Now()`.
  - Use `workflow.Sleep()` as a replacement for `time.Sleep()`.
- Working directly with threads or goroutines.
	- Use `workflow.Go()` as a replacement for the `go` statement.
    - Use `workflow.Channel()` as a replacement for the native `chan` type.
	Temporal provides support for both buffered and unbuffered channels.
	- Use `workflow.Selector()` as a replacement for the `select` statement.
- Iterating over data structures with unknown ordering.
  This includes iterating over maps using `range`, because with `range` the order of the map's iteration is randomized.
  Instead you can collect the keys of the map, sort them, and then iterate over the sorted keys to access the map.
  This technique provides deterministic results.
  You can also use a Side Effect or an Activity to process the map instead.
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

/*
If you run the BackgroundCheckNonDeterministic Workflow, eventually you will see a Workflow Task failure.

The Worker logs will show something similar to the following:

```shell
Error Potential deadlock detected: workflow goroutine "root" didn't yield for over a second StackTrace process event for backgroundcheck-replay-task-queue-local [panic]:
go.temporal.io/sdk/internal.(*coroutineState).call(0x1400001a780, 0x3b9aca00) ...
```

And you will see information about the failure in the Web UI as well.

![Web UI view of a non-determinism error](/img/non-deterministic-workflow-task-failure.png)

To inspect the Workflow Task failure using the Temporal CLI, you can use the `long` value for the `--fields` command option with the `temporal workflow show` command.

```shell
temporal workflow show \
 --workflow-id backgroundcheck_workflow_break \
 --namespace backgroundcheck_namespace \
 --fields long
```

This will display output similar to the following:

```shell
Progress:
  ID          Time                     Type                                                        Details
   1  2023-10-30T18:35:32Z  WorkflowExecutionStarted    {WorkflowType:{Name:BackgroundCheckNonDeterministic},
                                                        ParentInitiatedEventId:0,
                                                        TaskQueue:{Name:backgroundcheck-replay-task-queue-local,
                                                        Kind:Normal}, Input:["555-55-5555"],
                                                        WorkflowExecutionTimeout:0s, WorkflowRunTimeout:0s,
                                                        WorkflowTaskTimeout:10s, Initiator:Unspecified,
                                                        OriginalExecutionRunId:bf30d430-cc5c-445f-ad7b-e9e4a5cb9207,
                                                        Identity:temporal-cli:flossypurse@flossypurse-macbook-pro.local,
                                                        FirstExecutionRunId:bf30d430-cc5c-445f-ad7b-e9e4a5cb9207,
                                                        Attempt:1, FirstWorkflowTaskBackoff:0s,
                                                        ParentInitiatedEventVersion:0}
   2  2023-10-30T18:35:32Z  WorkflowTaskScheduled       {TaskQueue:{Name:backgroundcheck-replay-task-queue-local,
                                                        Kind:Normal}, StartToCloseTimeout:10s, Attempt:1}
   3  2023-10-30T18:35:32Z  WorkflowTaskStarted         {ScheduledEventId:2,
                                                        Identity:47041@flossypurse-macbook-pro.local@,
                                                        RequestId:6da86e56-cb43-4344-a138-019657e1d913,
                                                        SuggestContinueAsNew:false,
                                                        HistorySizeBytes:762}
   4  2023-10-30T18:35:33Z  WorkflowTaskFailed          {ScheduledEventId:2, StartedEventId:3, Cause:WorkflowWorkerUnhandledFailure,
                                                        Failure:{Message:Potential deadlock detected: workflow goroutine
                                                        "root" didn't yield for over a second, Source:GoSDK, StackTrace:process
                                                        event for backgroundcheck-replay-task-queue-local [panic]:
                                                        go.temporal.io/sdk/internal.(*coroutineState).call(0x1400061a780, 0x3b9aca00)
                                                                /Users/flossypurse/go/pkg/mod/go.temporal.io/sdk@v1.25.1/internal/internal_workflow.go:1011
                                                        +0x170 go.te ... poral.io/sdk@v1.25.1/internal/internal_worker_base.go:356 +0x48
                                                        created by go.temporal.io/sdk/internal.(*baseWorker).processTaskAsync in goroutine 15
                                                                /Users/flossypurse/go/pkg/mod/go.temporal.io/sdk@v1.25.1/internal/internal_worker_base.go:352
                                                        +0xbc, FailureInfo:{ApplicationFailureInfo: {Type:PanicError, NonRetryable:true}}},
                                                        Identity:47041@flossypurse-macbook-pro.local@, ForkEventVersion:0,
                                                        BinaryChecksum:48fa2bc5191e2e60e3f72a7a78d0e721}
```
*/

/* @dacx
id: backgroundcheck-replay-intrinsic-non-determinism
title: Intrinsic non-deterministic logic
description: This kind of logic prevents the Workflow code from executing to completion because the Workflow can take a different code path than the one expected from the Event History.
label: intrinsic-non-deterministic-logic
lines: 3-56
tags:
- tests
- replay
- event history
@dacx */

/* @dacx
id: backgroundcheck-replay-inspecting-the-non-deterministic-error
title: Intrinsic non-deterministic logic
description: This kind of logic prevents the Workflow code from executing to completion because the Workflow can take a different code path than the one expected from the Event History.
label: intrinsic-non-deterministic-logic
lines: 58-117
tags:
- tests
- replay
- event history
@dacx */
