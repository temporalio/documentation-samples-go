package backgroundcheckboilerplate

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

/*
Use the Temporal Go SDK's `testsuite` package to begin defining tests for Workflow and Activity code.
*/

const ssn string = "555-55-5555"

// Test_Workflow tests the BackgroundCheck Workflow
func Test_BackgroundCheck(t *testing.T) {
	// create a Workflow Test Suite
	testSuite := &testsuite.WorkflowTestSuite{}
	// Create a test environment
	env := testSuite.NewTestWorkflowEnvironment()
	// Mock the Activity Execution for the Workflow
	ssnTraceResult := "pass"
	env.OnActivity(SSNTraceActivity, mock.Anything, ssn).Return(&ssnTraceResult, nil)
	// Run the Workflow in the test environment
	env.ExecuteWorkflow(BackgroundCheck, ssn)
	// Check that the Workflow completed
	require.True(t, env.IsWorkflowCompleted())
	// Check there was no error returned
	require.NoError(t, env.GetWorkflowError())
	// Check for the expected value of the Workflow result
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))
}

// Test_SSNTraceActivity
func Test_SSNTraceActivity(t *testing.T) {
	//  Create a Workflow Test Suite
	testSuite := &testsuite.WorkflowTestSuite{}
	// Create a test environment
	env := testSuite.NewTestActivityEnvironment()
	// Register Activity with the enviroment
	env.RegisterActivity(SSNTraceActivity)
	// Run the Activity in the test enviroment
	val, err := env.ExecuteActivity(SSNTraceActivity, ssn)
	// Check there was no error on the call to execute the Activity
	require.NoError(t, err)
	// Check  that there was no error returned from the Activity
	var res string
	require.NoError(t, val.Get(&res))
	// Check for the expected return value.
	require.Equal(t, "pass", res)
}
