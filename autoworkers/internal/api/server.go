package api

import (
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"net/http"
	"autoworkers/internal/database"

)

type ApiServer struct{
	apistore *store.Store
	apiqueue *queue.Queue
	count int
	apidatabase *database.Database
}

func Constructor(q *queue.Queue, s *store.Store, d *database.Database) *ApiServer{
p := &ApiServer{
	apistore: s,
	apiqueue: q,
	count: 0,
	apidatabase: d,
}
return p
}


func (a *ApiServer) Start(){
	http.HandleFunc("/jobs", a.SubmitJob)
	http.HandleFunc("/jobs/",a.GetJob)
	http.ListenAndServe(":8080",nil)
}