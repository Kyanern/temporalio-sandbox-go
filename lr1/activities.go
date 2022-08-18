package lr1

import (
	"context"
	"math/rand"
	"strconv"

	"go.temporal.io/sdk/activity"
)

func Activity_randLR(ctx context.Context) (int, error) {
	logger := activity.GetLogger(ctx)
	r := rand.Int() % 2
	str := "Activity_randLR produced [" + strconv.Itoa(r) + "]"
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
