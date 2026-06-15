package api

import (
	"autoworkers/internal/job"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)
type SubmitJobRequest struct{
	Type string `json:"type"`
	Payload string `json:"payload"`
}
type SubmitJobResponse struct{
	ID string `json:"id"`
	Status job.JobStatus `json:"status"`
	Result string `json:"result"`
}

func (a *ApiServer) SubmitJob(w http.ResponseWriter, r *http.Request){
	if r.Method !="POST"{
		fmt.Fprint(w,"Route method is incorrect")
	}else{
		var b SubmitJobRequest
		x := json.NewDecoder(r.Body).Decode(&b)
		fmt.Println(x)
		a.count++
		testjob := &job.Job{
			ID: fmt.Sprintf("job-%d",a.count),
			Type: b.Type,
			Payload: b.Payload,
			Status: job.Pending,
		}
		store.Create(testjob,a.apistore)
		a.apidatabase.SaveJob(testjob)
		response := &SubmitJobResponse{
			ID : testjob.ID,
			Status : testjob.Status,
		}
		queue.Enqueue(a.apiqueue,testjob)
		json.NewEncoder(w).Encode(response)
	}

}

func  (a *ApiServer) GetJob(w http.ResponseWriter, r *http.Request){

	s := r.URL.Path
	jobid := strings.Split(s,"/")
	job := store.Get(a.apistore,jobid[2])
	if job==nil{
		fmt.Fprintln(w,"No jobs found")
	}else{
		json.NewEncoder(w).Encode(job)
	}
}