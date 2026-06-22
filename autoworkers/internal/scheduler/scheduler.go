package scheduler

import (
	"autoworkers/internal/redis"
	"fmt"
	"time"
)


type Scheduler struct{
	redis *redis.Redis
}

func Constructor(r *redis.Redis)*Scheduler{
	return &Scheduler{
		redis: r,
	}
}

func (s *Scheduler)Start(){
for {
	now := time.Now().Unix()
	fmt.Println("scheduler tick",now)
	time.Sleep(1 * time.Second)
}
}