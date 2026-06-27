package llm

import "autoworkers/internal/job"

type Client interface{
	Generate(j *job.Job) (string,error)
}

