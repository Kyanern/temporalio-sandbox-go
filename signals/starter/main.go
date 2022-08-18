package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"go.temporal.io/sdk/client"

	"github.com/google/uuid"
	"github.com/kyanern/temporalio-sandbox-go/signals"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "sb_signals_" + uuid.NewString(),
		TaskQueue: "sb_signals",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, signals.Workflow_signals)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	timeseed := time.Since(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Now().Location()))
	rand.Seed(timeseed.Milliseconds())
	choice := rand.Int() % 2
	if choice == 1 {
		log.Println("Sleeping for 3 seconds before sending continue signal...")
		time.Sleep(3 * time.Second)
		sig := signals.ContinueSignal{Command: signals.CommandTypes.CONTINUE}
		err := c.SignalWorkflow(context.Background(), workflowOptions.ID, "", signals.SignalChannels.CONTINUE_CHANNEL, sig)
		if err != nil {
			log.Fatalln("unable to signal continue to workflow", err)
		}
	} else {
		log.Println("Sleeping for 3 seconds before sending cancel signal...")
		time.Sleep(3 * time.Second)
		sig := signals.CancelSignal{Command: signals.CommandTypes.CANCEL}
		err := c.SignalWorkflow(context.Background(), workflowOptions.ID, "", signals.SignalChannels.CANCEL_CHANNEL, sig)
		if err != nil {
			log.Fatalln("unable to signal cancel to workflow", err)
		}
	}

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
