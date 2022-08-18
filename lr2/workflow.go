package lr2

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// Workflow_LR2 is a workflow definition.
// It decides whether to go left or right
// using the number given to it by its caller.
func Workflow_LR2(ctx workflow.Context, i int) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("sb_lr2 workflow started")

	var choiceresult int
	err := workflow.ExecuteActivity(ctx, Activity_chooseLR, i).Get(ctx, &choiceresult)
	if err != nil {
		logger.Error("Activity_chooseLR failed.", "Error", err)
		return "", err
	}

	const LEFT = 0
	const RIGHT = 1
	var result string
	switch choiceresult {
	case LEFT:
		workflow.ExecuteActivity(ctx, Activity_Left).Get(ctx, &result)
	case RIGHT:
		workflow.ExecuteActivity(ctx, Activity_Right).Get(ctx, &result)
	default:
		result = "Didn't go either LEFT or RIGHT."
	}

	return result, nil
}
