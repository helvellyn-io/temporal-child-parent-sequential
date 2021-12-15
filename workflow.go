package temporalweather

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GetWeatherWorkflow(ctx workflow.Context) (float64, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result float64
	err := workflow.ExecuteActivity(ctx, GetWeather).Get(ctx, &result)
	return result, err
}
