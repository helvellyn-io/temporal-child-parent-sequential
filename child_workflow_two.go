package temporalweather

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

//Gets Temps
func ChildWorkflowTwo(ctx workflow.Context) (float64, error) {
	logger := workflow.GetLogger(ctx)

	options := workflow.ActivityOptions{
		ScheduleToStartTimeout: 10 * time.Second,
		StartToCloseTimeout:    5 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
		HeartbeatTimeout:       0,
	}
	var result float64
	ctx = workflow.WithActivityOptions(ctx, options)
	err := workflow.ExecuteActivity(ctx, GetHumidity).Get(ctx, &result)
	logger.Info("2nd Child Workflow execution.")
	return result, err
}
