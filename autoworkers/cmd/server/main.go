package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/database"
	"autoworkers/internal/executor"
	"autoworkers/internal/llm"
	"autoworkers/internal/manager"
	"autoworkers/internal/metrics"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
	}
	s := store.Constructor()
	d := database.Constructor()
	r := redis.Constructor()
	m := metrics.Constructor()
	llmClient := llm.Constructor(os.Getenv("OPENROUTER_API_KEY"))
	exec := executor.Constructor(llmClient)
	ma := manager.Constructor(r,s,d,m,exec)
	server := api.Constructor(r,s,d,m)
	go ma.Start()
	server.Start()

}