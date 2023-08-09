package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-gin/pkg/setting"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func Setup() {
	// 官方文档 https://redis.uptrace.dev/zh/
	ctx := context.Background()
	hostStr := setting.RedisSetting.Host
	nodes := strings.Split(hostStr, ",")
	fmt.Println(nodes)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        nodes,
		Password:     setting.RedisSetting.Password,
		MaxIdleConns: setting.RedisSetting.MaxIdle,
	})
	fmt.Println(rdb)
	rdb.SetNX(ctx, "test", "a", 1*time.Minute)
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
