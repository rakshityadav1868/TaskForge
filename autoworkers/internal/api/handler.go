package api

import (
	"autoworkers/internal/job"
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
type MetricsResponse struct {
    Pending     int `json:"pending"`
    Running     int `json:"running"`
    Completed   int `json:"completed"`
    Failed      int `json:"failed"`
}

func (a *ApiServer) SubmitJob(w http.ResponseWriter, r *http.Request){

	
	if r.Method !="POST"{
		fmt.Fprint(w,"Route method is incorrect")
	}else{
		var b SubmitJobRequest
		err := json.NewDecoder(r.Body).Decode(&b)
		if err!=nil{
			fmt.Println(err)
		}
		a.count++
		testjob := &job.Job{
			ID: fmt.Sprintf("job-%d",a.count),
			Type: b.Type,
			Payload: b.Payload,
			Status: job.Pending,
			RetryCount: 0,
			MaxRetries: 3,
		}
		store.Create(testjob,a.apistore)
		a.apidatabase.SaveJob(testjob)
		a.apimetrices.Pending++
		response := &SubmitJobResponse{
			ID : testjob.ID,
			Status : testjob.Status,
		}
		a.apiredis.Enqueue(testjob.ID)
		json.NewEncoder(w).Encode(response)
	}

}

func  (a *ApiServer) GetJob(w http.ResponseWriter, r *http.Request){

	s := r.URL.Path
	jobid := strings.Split(s,"/")
	job := a.apidatabase.GetJob(jobid[2])
	if job==nil{
		fmt.Fprintln(w,"No jobs found")
	}else{
		json.NewEncoder(w).Encode(job)
	}
}

func (a *ApiServer) GetAllJobs(w http.ResponseWriter, r *http.Request){
	x := a.apidatabase.GetAllJobs()
	json.NewEncoder(w).Encode(x)
}
func (a *ApiServer) GetMetrics(w http.ResponseWriter, r *http.Request){
	response := MetricsResponse{
		Pending: a.apimetrices.Pending,
		Running: a.apimetrices.Running,
		Completed: a.apimetrices.Completed,
		Failed: a.apimetrices.Failed,
	}
	json.NewEncoder(w).Encode(response)
}