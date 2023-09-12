package backgroundcheckboilerplate

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

/*
We recommend organizing Workflow code together with other Workflow code.
For example, in a small project like this, it is still a best practice to have a dedicated file for the Workflow.
*/

// BackgroundCheck is your custom Workflow Definition.
func BackgroundCheck(ctx workflow.Context, param string) (string, error) {
	// Define the Activity Execution options
	// StartToCloseTimeout or ScheduleToCloseTimeout must be set
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	// Execute the Activity synchronously (wait for the result before proceeding)
	var ssnTraceResult string
	err := workflow.ExecuteActivity(ctx, SSNTraceActivity, param).Get(ctx, &ssnTraceResult)
	if err != nil {
		return "", err
	}
	// Make the results of the Workflow available
	return ssnTraceResult, nil
}
