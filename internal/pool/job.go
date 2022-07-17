package pool

import (
	"context"
)

type Job interface {
	Execute(ctx context.Context) JobResult
}
