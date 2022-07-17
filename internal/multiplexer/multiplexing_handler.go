package multiplexer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/morfin60/parallel-handlers/internal/pool"
)

type Data struct {
	Code  int    `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// Handler to check urls
func MultiplexingHandler(poolSize int32) http.HandlerFunc {
	workerPools := &sync.Pool{}

	return func(w http.ResponseWriter, r *http.Request) {
		var workerPool *pool.Pool

		if wp := workerPools.Get(); wp == nil {
			workerPool = pool.New(poolSize)
		} else {
			workerPool = wp.(*pool.Pool)
		}
		ctx, cancel := context.WithCancel(context.Background())
		response := &http.Response{StatusCode: http.StatusOK}
		responseData := &Data{Code: http.StatusOK}

		defer func() {
			cancel()

			data, err := json.Marshal(responseData)
			if err != nil {
				log.Printf("Failed to serialize value: %#v error: %s", data, err.Error())
			} else {
				response.Body = io.NopCloser(bytes.NewReader(data))
			}

			response.Write(w)
		}()

		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responseData.Code = http.StatusBadRequest
		}

		urls := []string{}
		err = json.Unmarshal(requestData, &urls)
		if err != nil {
			responseData.Code = http.StatusBadRequest
			responseData.Error = "Failed to process request"

			return
		}

		// No urls received in request
		if len(urls) == 0 {
			responseData.Data = []struct{}{}

			return
		}

		// Too many urls
		if len(urls) > 20 {
			responseData.Code = http.StatusBadRequest
			responseData.Error = "Too many urls"

			return
		}

		resultsList := make([]*Result, 0, len(urls))

		jobs := make([]pool.Job, len(urls))

		for i, url := range urls {
			jobs[i] = &MultiplexerUrlJob{url: url}
		}

		workerPool.Start(ctx)
		workerPool.ProcessJobs(jobs)

		for jobResult := range workerPool.Results() {
			if jobResult.Err != nil {
				responseData.Code = 500
				responseData.Error = "Failed to process url " + jobResult.Job.(*MultiplexerUrlJob).url

				return
			} else {
				result := jobResult.Value.(*Result)
				resultsList = append(resultsList, result)
			}
		}

		responseData.Data = resultsList
	}
}
