package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/kyanern/temporalio-sandbox-go/lr1"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "sb_lr1", worker.Options{})

	w.RegisterWorkflow(lr1.Workflow_LR1)
	w.RegisterActivity(lr1.Activity_randLR)
	w.RegisterActivity(lr1.Activity_Left)
	w.RegisterActivity(lr1.Activity_Right)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
