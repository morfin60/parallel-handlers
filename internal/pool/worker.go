package pool

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- JobResult) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job.Execute(ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- JobResult{
				Err: ctx.Err(),
			}
			return
		}
	}
}
