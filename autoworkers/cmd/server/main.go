package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
	"fmt"
	"autoworkers/internal/database"
)

func main(){
	s := store.Constructor()
	q := queue.Constructor()
	w :=worker.Constructor(q, s)
	d := database.Constructor()
	fmt.Println(d)
	server := api.Constructor(q,s,d)
	go worker.Workers(w)
	server.Start()

}