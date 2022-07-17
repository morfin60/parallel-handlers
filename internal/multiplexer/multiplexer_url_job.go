package multiplexer

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/morfin60/parallel-handlers/internal/pool"
)

type MultiplexerUrlJob struct {
	url string
}

func NewUrlJob(url string) *MultiplexerUrlJob {
	return &MultiplexerUrlJob{url: url}
}

func (muj *MultiplexerUrlJob) Execute(ctx context.Context) pool.JobResult {
	requestCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	client := &http.Client{}

	// Perform request with context created from parent with timeout
	req, err := http.NewRequest("GET", muj.url, nil)
	req = req.WithContext(requestCtx)
	resp, err := client.Do(req)
	result := &Result{Success: false}
	jobResult := pool.JobResult{Job: muj}

	if err != nil {
		jobResult.Err = err
	} else {
		result.Url = muj.url
		result.Code = resp.StatusCode

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			jobResult.Err = err
		} else {
			result.Success = true
			result.Content = content
		}

		jobResult.Value = result
	}

	return jobResult
}
