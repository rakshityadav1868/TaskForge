package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/database"
	"autoworkers/internal/manager"
	"autoworkers/internal/metrics"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
)

func main(){
	s := store.Constructor()
	d := database.Constructor()
	r := redis.Constructor()
	m := metrics.Constructor()
	ma := manager.Constructor(r,s,d,m)
	server := api.Constructor(r,s,d,m)
	go ma.Start()
	server.Start()

}