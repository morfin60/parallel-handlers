package pool

type JobResult struct {
	Value any   `json:"value"`
	Err   error `json:"error"`
	Job   Job
}
