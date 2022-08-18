package main

import (
	"context"
	"log"
	"math/rand"

	"go.temporal.io/sdk/client"

	"github.com/google/uuid"
	"github.com/kyanern/temporalio-sandbox-go/lr2"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "sb_lr2_" + uuid.NewString(),
		TaskQueue: "sb_lr2",
	}

	n := rand.Int()
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, lr2.Workflow_LR2, n)
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
