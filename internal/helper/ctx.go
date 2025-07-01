package helper

import (
	"context"
	"time"
)

func CreateCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return  ctx, cancel
}