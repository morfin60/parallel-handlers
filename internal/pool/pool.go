package pool

import (
	"context"
	"sync"
)

type Pool struct {
	workersNumber int32
	jobs          chan Job
	results       chan JobResult
	Done          chan struct{}
}

func New(workersNumber int32) *Pool {
	return &Pool{
		workersNumber: workersNumber,
		jobs:          make(chan Job, workersNumber),
		results:       make(chan JobResult, workersNumber),
	}
}

func (p *Pool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	for i := int32(0); i < p.workersNumber; i++ {
		wg.Add(1)
		go worker(ctx, &wg, p.jobs, p.results)
	}

	go func() {
		defer close(p.results)
		wg.Wait()
	}()
}

// Get results of last operation
func (p Pool) Results() <-chan JobResult {
	return p.results
}

// Process jobs from slice
func (p Pool) ProcessJobs(jobs []Job) {
	for i, _ := range jobs {
		p.jobs <- jobs[i]
	}

	close(p.jobs)
}
