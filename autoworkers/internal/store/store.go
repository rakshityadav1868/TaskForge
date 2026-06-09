package store

import (
	"autoworkers/internal/job"
	"fmt"
	"sync"
)

type Store struct{
	jobmapping map[string] *job.Job
	mu sync.RWMutex
}

func Constructor() *Store{
makejob := make(map[string] *job.Job)
p:=  &Store{
	jobmapping: makejob,
}
return p
}
func Create(x *job.Job, p *Store){
	p.mu.Lock()

	defer p.mu.Unlock()
	p.jobmapping[x.ID] = x
}

func Get(p *Store, jobId string) *job.Job {
	p.mu.RLock()
	defer p.mu.RUnlock()
	_,ok := p.jobmapping[jobId]
	y := p.jobmapping[jobId]

	if (ok){
		fmt.Println("Job found")
		return y
		
	}else{
		fmt.Println("Job not found")
		return nil
	}

}

func UpdateStatus(x *job.Job,p *Store){
	p.mu.Lock()
	defer p.mu.Unlock()
	jobFound,ok :=p.jobmapping[x.ID]
	if ok{
		jobFound.Status=x.Status
		jobFound.Created_time=x.Created_time
		jobFound.Finished_time=x.Finished_time


	}else{
		fmt.Println("job not found")
		
	}
}
func main(){


}