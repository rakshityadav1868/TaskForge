package worker

import (
	"autoworkers/internal/database"
	"autoworkers/internal/executor"
	"autoworkers/internal/job"
	"autoworkers/internal/metrics"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"fmt"
)
type Worker struct{
	id int
	redisqueue *redis.Redis
	store *store.Store
	database *database.Database
	metrics *metrics.Metrics
}
func Constructor(id int ,redisqueue *redis.Redis,store *store.Store, database *database.Database, metrics *metrics.Metrics) *Worker{
	s := &Worker{
		id: id,
		redisqueue: redisqueue,
		store: store,
		database: database,
		metrics: metrics,
	}
	return s

}

func Workers(m *Worker){
	for{
		jobId := m.redisqueue.Dequeue()
		fmt.Printf("Worker %d processing %s\n",m.id,jobId)
		jobobj := store.Get(m.store,jobId)
		if jobobj==nil{
			fmt.Println("No job found")
			continue
		}else{
			oldstatus := jobobj.Status
			jobobj.Status = job.Running
			m.metrics.Update(oldstatus,jobobj.Status)
			
			store.UpdateStatus(jobobj,m.store)
			m.database.UpdateJob(jobobj)
			result ,err := executor.Execute(jobobj)
			if err!=nil{
				jobobj.RetryCount ++
				if jobobj.RetryCount<jobobj.MaxRetries{
					oldstatus := jobobj.Status
					jobobj.Status = job.Pending
					m.metrics.Update(oldstatus,jobobj.Status)
					store.UpdateStatus(jobobj,m.store)
					m.database.UpdateJob(jobobj)
					// retry count only when error appear
					fmt.Printf( "Retrying %s (%d/%d)\n",jobobj.ID,jobobj.RetryCount,jobobj.MaxRetries)
					m.redisqueue.Enqueue(jobId)	
					continue
				}else{
					// permanent failure log
					fmt.Printf( "Job %s permanently failed\n",jobobj.ID)
					oldstatus := jobobj.Status
					jobobj.Status = job.Failed
					m.metrics.Update(oldstatus,jobobj.Status)
					Error := err.Error()
					jobobj.Error = Error
					store.UpdateStatus(jobobj,m.store)
					m.database.UpdateJob(jobobj)
					continue
				}
			}else{
				jobobj.Error = ""
				jobobj.Result = result
				oldstatus := jobobj.Status
				jobobj.Status = job.Completed
				m.metrics.Update(oldstatus,jobobj.Status)
				store.UpdateStatus(jobobj,m.store)
				m.database.UpdateJob(jobobj)
			}
		}
	}

}