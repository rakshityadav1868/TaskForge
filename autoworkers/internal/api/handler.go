package api

import (
	"autoworkers/internal/job"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"fmt"
	"net/http"
	"strings"
)

func (a *ApiServer) SubmitJob(w http.ResponseWriter, r *http.Request){
	fmt.Println(a.apistore)
	fmt.Println(a.apiqueue)
	a.count++
	testjob := &job.Job{
		ID: fmt.Sprintf("job-%d",a.count),
		Type: "test",
		Payload: "hello world",
		Status: job.Pending,
	}
	store.Create(testjob,a.apistore)
	queue.Enqueue(a.apiqueue,testjob)
	fmt.Fprintln(w,testjob.Status)
	fmt.Fprintln(w,testjob.ID)

}

func  (a *ApiServer) GetJob(w http.ResponseWriter, r *http.Request){
	s := r.URL.Path
	jobid := strings.Split(s,"/")
	job := store.Get(a.apistore,jobid[2])
	if job==nil{
		fmt.Fprintln(w,"No jobs found")
	}else{
		fmt.Fprint(w,job.ID, "\n")
		fmt.Fprint(w,job.Status, "\n")
		fmt.Fprint(w,job.Result, "\n")
	}
}