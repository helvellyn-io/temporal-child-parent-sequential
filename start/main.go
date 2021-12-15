package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	app "temporal-weather/app"
)

func main() {
	// Client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "weather-workflow",
		TaskQueue: app.WeatherTaskQueue,
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GetWeatherWorkflow)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	fmt.Println(we)
}
