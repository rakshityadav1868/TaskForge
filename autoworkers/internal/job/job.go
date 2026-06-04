package job


type JobStatus int

const (
	Pending JobStatus=iota
	Running
	Completed
	failed
)

type Job struct{
	ID string
	Tyype string 
	Payload string
	Status string
	Result string
	Error string
	Created_time int
	Started_time int
	Finished_time int 

}

func main (){

}