package demo_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
//func initClient() (err error) {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	_, err = rdb.Ping().Result()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func DemoRedis() {
	V8Example()
	//err := initClientV8()
	//if err != nil {
	//	fmt.Printf("连接失败::%v", err)
	//} else {
	//	fmt.Print("连接成功")
	//}
}

// 初始化连接
func initClientV8() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClientV8(); err != nil {
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	err2 := rdb.Set(ctx, "key2", "我是Key2的值", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
