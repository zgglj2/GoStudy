package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// redisdb := redis.NewFailoverClient(&redis.FailoverOptions{
	// 	MasterName:    "master",
	// 	SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	// })

	// redisdb := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	// })

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func redisGetSet() {
	err := redisdb.Set("score", 10, 0).Err()
	if err != nil {
		fmt.Println("set score failed, err: ", err)
		return
	}
	val, err := redisdb.Get("score").Result()
	if err != nil {
		fmt.Println("get score failed, err: ", err)
		return
	}
	fmt.Println("score: ", val)

	val2, err := redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
		return
	} else if err != nil {
		fmt.Println("get name failed, err: ", err)
		return
	} else {
		fmt.Println("name: ", val2)
	}
}

func redisZGetSet() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	num, err := redisdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d record success\n", num)

	// 把Golang的分数加10
	newScore, err := redisdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := redisdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func redisPipeline() {
	pipe := redisdb.Pipeline()

	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func redisPipelined() {
	var incr *redis.IntCmd
	_, err := redisdb.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipelined_counter")
		pipe.Expire("pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr.Val(), err)
}

func redisTxPipeline() {
	pipe := redisdb.TxPipeline()

	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func redisTxPipelined() {
	var incr *redis.IntCmd
	_, err := redisdb.TxPipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("tx_pipelined_counter")
		pipe.Expire("tx_pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr.Val(), err)
}

func redisWatch() {
	// 监视watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	redisdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("connect to redis failed, err: ", err)
		return
	}
	fmt.Println("connect to redis success")

	redisGetSet()
	redisZGetSet()

	redisPipeline()
	redisPipelined()

	redisTxPipeline()
	redisTxPipelined()

	redisWatch()
}
