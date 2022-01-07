package main

import (
	"log"

	app "temporal-weather/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions
	w := worker.New(c, "child-workflow", worker.Options{})

	//w := worker.New(c, app.WeatherTaskQueue, worker.Options{})
	//w.RegisterWorkflow(app.ParentWorkflow)
	w.RegisterWorkflow(app.ChildWorkflowOne)
	w.RegisterWorkflow(app.ChildWorkflowTwo)
	w.RegisterActivity(app.GetHumidity)
	w.RegisterActivity(app.GetTemp)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}

}
