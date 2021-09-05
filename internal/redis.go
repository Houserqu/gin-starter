package internal

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisCtx = context.Background()
var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:55000",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := Redis.Ping(RedisCtx).Result(); err != nil {
		log.Fatal(err.Error())
		return
	}

	//err := rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
}
