package signals

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

//Workflow_signals is a workflow definition.
//It has an activity that waits for a signal before
//execution continues.
func Workflow_signals(ctx workflow.Context) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var retval string
	const FUNCNAME = "Workflow_signals"
	log.Println("Entered " + FUNCNAME + "...")

	logger := workflow.GetLogger(ctx)

	var result string
	workflow.ExecuteActivity(ctx, Activity_stub1).Get(ctx, &result)
	logger.Info(result)
	log.Println(result)

	continueChannel := workflow.GetSignalChannel(ctx, SignalChannels.CONTINUE_CHANNEL)
	cancelChannel := workflow.GetSignalChannel(ctx, SignalChannels.CANCEL_CHANNEL)
	cancelled := false

	selector := workflow.NewSelector(ctx)
	selector.AddReceive(continueChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal interface{}
		c.Receive(ctx, &signal)

		//workflow execution shall continue. Activity_stub2 should be run.
	})
	selector.AddReceive(cancelChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal interface{}
		c.Receive(ctx, &signal)
		cancelled = true
		//workflow execution shall stop. Activity_stub2 should not be run.
	})

	selector.Select(ctx)

	if cancelled {
		log.Println("Leaving " + FUNCNAME + ". Returning a string.")
		retval = FUNCNAME + " got cancelled."
		return retval, nil
	}

	workflow.ExecuteActivity(ctx, Activity_stub2).Get(ctx, &result)
	logger.Info(result)
	log.Println(result)

	log.Println("Leaving " + FUNCNAME + ". Returning a string.")

	retval = FUNCNAME + " went through all the way."

	return retval, nil
}
