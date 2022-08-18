package lr1

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// Workflow_LR1 is a workflow definition.
// It decides whether to go left or right
// using a pseudo-random number it generated.
func Workflow_LR1(ctx workflow.Context) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("sb_lr1 workflow started")

	var randresult int
	err := workflow.ExecuteActivity(ctx, Activity_randLR).Get(ctx, &randresult)
	if err != nil {
		logger.Error("Activity_randLR failed.", "Error", err)
		return "", err
	}

	const LEFT = 0
	const RIGHT = 1
	var result string
	switch randresult {
	case LEFT:
		workflow.ExecuteActivity(ctx, Activity_Left).Get(ctx, &result)
	case RIGHT:
		workflow.ExecuteActivity(ctx, Activity_Right).Get(ctx, &result)
	default:
		result = "Didn't go either LEFT or RIGHT."
	}

	return result, nil
}
