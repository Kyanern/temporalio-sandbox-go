package lr2

import (
	"context"
	"strconv"

	"go.temporal.io/sdk/activity"
)

func Activity_chooseLR(ctx context.Context, i int) (int, error) {
	logger := activity.GetLogger(ctx)
	r := i % 2
	str := "Activity_chooseLR produced [" + strconv.Itoa(r) + "]"
	logger.Info(str)
	return r, nil
}

func Activity_Left(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	r := "Activity_Left"
	logger.Info(r)
	return r, nil
}

func Activity_Right(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	r := "Activity_Right"
	logger.Info(r)
	return r, nil
}
