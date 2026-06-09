package executor

import (
	"autoworkers/internal/job"
	"time"
)

func Execute(j *job.Job) string{
	time.Sleep(2 * time.Second)
	return  "success"
}
