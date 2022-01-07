package temporalweather

import (
	"fmt"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/workflow"
)

//Schedules the child activities
func ParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	workflowID := "parent-workflow_" + uuid.New()

	options := workflow.ChildWorkflowOptions{
		WorkflowID: workflowID,
	}

	ctx = workflow.WithChildOptions(ctx, options)
	//run child workflow 1
	var result_1 float64
	err := workflow.ExecuteChildWorkflow(ctx, ChildWorkflowOne).Get(ctx, &result_1)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)

	}
	//run child workflow 2
	var result_2 int
	err = workflow.ExecuteChildWorkflow(ctx, ChildWorkflowTwo).Get(ctx, &result_2)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", err)
	}
	logger.Info("Parent execution completed.")

	a := fmt.Sprint(result_1, result_2, err)
	return a, err

}
