package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/kyanern/temporalio-sandbox-go/signals"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "sb_signals", worker.Options{})

	w.RegisterWorkflow(signals.Workflow_signals)
	w.RegisterActivity(signals.Activity_stub1)
	w.RegisterActivity(signals.Activity_stub2)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
