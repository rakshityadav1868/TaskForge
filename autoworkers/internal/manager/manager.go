package manager

import (
	"autoworkers/internal/database"
	"autoworkers/internal/metrics"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
	"fmt"
	"math"
	"time"
)

type Manager struct{
	redisqueue *redis.Redis
	store *store.Store
	database *database.Database
	metrics *metrics.Metrics
	workercount int
	workers map[int]*worker.Worker
}

func Constructor(redisqueue *redis.Redis,store *store.Store, database *database.Database, metrics *metrics.Metrics) *Manager{
return &Manager{
	redisqueue: redisqueue,
	store: store,
	database: database,
	metrics: metrics,
	workers: make(map[int]*worker.Worker),

}
}

func (m *Manager) StartWorker(){
	m.workercount  ++
	w1 := worker.Constructor(m.workercount,m.redisqueue,m.store,m.database,m.metrics)
	m.workers[m.workercount] = w1
	go worker.Workers(w1)
}

func (m *Manager) Start() {
	m.StartWorker()
	for {
		queuelength := m.redisqueue.GetQueueLength()
		neededworkers := m.CalculateWorkers(queuelength)
		if neededworkers>m.workercount{
			workerstostart:= neededworkers-m.workercount
			for i := range workerstostart{
				m.StartWorker()
			}
		}
		fmt.Println(queuelength)
		time.Sleep(1 * time.Second) 
	}
	
}
func (m *Manager) CalculateWorkers(queuelength int) int {
	maxworkers := 15
	neededworkers := math.Ceil(float64(queuelength)/5)
	if neededworkers < 1 {
    return 1
	}
	if neededworkers > float64(maxworkers){
		return maxworkers
	}else{

		return  int(neededworkers)
	}
}