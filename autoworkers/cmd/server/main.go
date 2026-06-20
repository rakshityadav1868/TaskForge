package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/database"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
)

func main(){
	s := store.Constructor()
	d := database.Constructor()
	r := redis.Constructor()
	w1 :=worker.Constructor(1,r,s,d)
	w2 :=worker.Constructor(2,r,s,d)
	w3 :=worker.Constructor(3,r,s,d)
	server := api.Constructor(r,s,d)
	go worker.Workers(w1)
	go worker.Workers(w2)
	go worker.Workers(w3)
	server.Start()

}