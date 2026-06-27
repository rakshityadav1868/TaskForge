package manager

import (
	"autoworkers/internal/database"
	"autoworkers/internal/executor"
	"autoworkers/internal/metrics"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
	"context"
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
	cancels map[int]context.CancelFunc
	executor *executor.Executor
}

func Constructor(redisqueue *redis.Redis,store *store.Store, database *database.Database, metrics *metrics.Metrics, exe *executor.Executor) *Manager{
return &Manager{
	redisqueue: redisqueue,
	store: store,
	database: database,
	metrics: metrics,
	workers: make(map[int]*worker.Worker),
	cancels: make(map[int]context.CancelFunc),
	executor: exe,

}
}

func (m *Manager) StartWorker(){
	m.workercount  ++
	ctx, cancel := context.WithCancel(context.Background())
	w1 := worker.Constructor(m.workercount,m.redisqueue,m.store,m.database,m.metrics,ctx,m.executor)
	m.workers[m.workercount] = w1
	m.cancels[m.workercount]=cancel
	go worker.Workers(w1)
}

func (m *Manager) Start() {
	m.StartWorker()
	for {
		queuelength := m.redisqueue.GetQueueLength()
		neededworkers := m.CalculateWorkers(queuelength)
		if neededworkers>m.workercount{
			workerstostart:= neededworkers-m.workercount
			for i := 0; i<workerstostart;i++{
				m.StartWorker()
			}
		}
		if neededworkers <m.workercount{
			workerstostop := m.workercount - neededworkers
			for i := 0; i < workerstostop; i++ {
        		m.StopWorker()
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

func (m *Manager) StopWorker() {
	if m.workercount <= 1 {
    return
	}	
    id := m.workercount
	m.cancels[id]()
	delete(m.workers,id)
	delete(m.cancels,id)
	m.workercount --

}