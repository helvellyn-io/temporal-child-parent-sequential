package temporalweather

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

//Gets Temps
func ChildWorkflowOne(ctx workflow.Context) (float64, error) {
	logger := workflow.GetLogger(ctx)

	options := workflow.ActivityOptions{
		ScheduleToStartTimeout: 10 * time.Second,
		StartToCloseTimeout:    5 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
		HeartbeatTimeout:       0,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var rresult float64
	err := workflow.ExecuteActivity(ctx, GetTemp).Get(ctx, &rresult)
	logger.Info("1st Child Workflow execution.")
	return rresult, err
}
