package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct{
	Client *goredis.Client
}
func Constructor() *Redis{
	ctx := context.Background()
	r := goredis.NewClient(&goredis.Options{Addr: "localhost:6379",
										Password: "",
										DB: 0})
	_,err := r.Ping(ctx).Result()
	if err!=nil{
		fmt.Println(err)

	}
	st := &Redis{
		Client: r,
	}
	return st
	
}


func (r *Redis) Enqueue(jobID string){
	ctx := context.Background()
	x := r.Client.RPush(ctx,"jobs",jobID)
	if x.Err()!=nil{
		fmt.Println(x.Err())
	}
}

func (r *Redis) Dequeue() string{
	ctx := context.Background()
	x := r.Client.BLPop(ctx,0,"jobs")
	if x.Err()!=nil{
		fmt.Println(x.Err())
	}else{
		result ,err := x.Result()
		if err!=nil{
			fmt.Println(err)
		}else{
			return result[1]
		}
	}
	return ""
}
func (r* Redis) EnqueueDelay(jobID string,executeAt int64,){
	ctx := context.Background()
	x := r.Client.ZAdd(
		ctx,
		"delayed_jobs",
		goredis.Z{
			Score: float64(executeAt),
			Member: jobID,
		},
	)
	if x.Err()!=nil{
		fmt.Println(x.Err())
	}

}