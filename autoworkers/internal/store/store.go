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

func Get(x *job.Job,p *Store) {
	p.mu.RLock()
	defer p.mu.Unlock()
	_,ok := p.jobmapping[x.ID]

	if (ok){

		fmt.Println("Job found")
	}else{
		fmt.Println("Job not found")
	}

}

func UpdateStatus(x *job.Job,p *Store){
	jobFound,ok :=p.jobmapping[x.ID]
	if ok{
		p.mu.Lock()
		jobFound.Status=x.Status
		jobFound.Created_time=x.Created_time
		jobFound.Finished_time=x.Finished_time
		p.mu.Unlock()

	}else{
		fmt.Println("job not found")
		
	}
}
func main(){


}