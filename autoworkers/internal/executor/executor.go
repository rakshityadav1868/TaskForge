package executor

import (
	"autoworkers/internal/job"
	"autoworkers/internal/llm"
	"errors"
	"time"
)


type Executor struct{
	llm llm.Client
}

func Constructor(client llm.Client) *Executor{
	return &Executor{
		llm: client,
	}
}

func (e *Executor) Execute(j *job.Job) (string,error){
	time.Sleep(2 * time.Second)
	if j.Type=="llm"{
		return e.llm.Generate(j)
	}
	if j.Type=="fail"{
		err := errors.New("failed")
		return "", err
	}else{
		return "success", nil
	}
}
