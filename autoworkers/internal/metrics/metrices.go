package metrics

import (
	"autoworkers/internal/job"
)


type Metrics struct {
    Pending   int
    Running   int
    Completed int
    Failed    int
}

func Constructor() *Metrics{
	return &Metrics{
		Pending: 0,
		Running: 0,
		Completed: 0,
		Failed: 0,
	}
}

func (m *Metrics) Update(oldStatus job.JobStatus,newStatus job.JobStatus){
	switch oldStatus{
	case job.Pending:
		m.Pending--
	case job.Running:
		m.Running--
	case job.Completed:
		m.Completed--
	case job.Failed:
		m.Failed--
	}
	switch newStatus{
	case job.Pending:
		m.Pending++
	case job.Running:
		m.Running++
	case job.Completed:
		m.Completed++
	case job.Failed:
		m.Failed++
	}

	
}