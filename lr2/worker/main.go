package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/kyanern/temporalio-sandbox-go/lr2"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "sb_lr2", worker.Options{})

	w.RegisterWorkflow(lr2.Workflow_LR2)
	w.RegisterActivity(lr2.Activity_chooseLR)
	w.RegisterActivity(lr2.Activity_Left)
	w.RegisterActivity(lr2.Activity_Right)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
