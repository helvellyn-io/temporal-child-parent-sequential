package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
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
	workflowID := "parent-workflow_" + uuid.New()

	//workflowID := "Starter-Workflow"
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "child-workflow",
	}
	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.ChildWorkflowOne)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	var result_1 float64
	err = workflowRun.Get(context.Background(), &result_1)
	if err != nil {
		log.Fatalln("Failure getting workflow result", err)
	}
	log.Println("Workflow result child workflow 1", result_1)

	//workflow two
	var result_2 int
	workflowRun, err = c.ExecuteWorkflow(context.Background(), workflowOptions, app.ChildWorkflowTwo)
	if err != nil {
		log.Fatalln("Unable to execture workflow", err)
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), &result_2)

	log.Println("Workflow result child workflow 2", result_2)

}
