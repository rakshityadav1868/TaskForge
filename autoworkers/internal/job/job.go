package job


type JobStatus int

const (
	Pending JobStatus=iota
	Running
	Completed
	Failed
)

type Job struct{
	ID string
	Type string 
	Payload string
	Status JobStatus
	Result string
	Error string
	Created_time int
	Started_time int
	Finished_time int 
	RetryCount int
	MaxRetries int

}