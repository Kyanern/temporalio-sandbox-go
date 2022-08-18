package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kyanern/temporalio-sandbox-go/lr1"
	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "sb_lr1_" + uuid.NewString(),
		TaskQueue: "sb_lr1",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, lr1.Workflow_LR1)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
