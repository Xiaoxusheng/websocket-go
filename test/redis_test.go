package test

import (
	"Gin/db"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	// 创建一个Redis客户端

	ctx := context.Background()
	//// 创建一个消费者组
	group := "1"
	stream := "my-stream1"
	consumerName := "my-consumer"
	r, err := db.Rdb.XGroupCreateMkStream(ctx, stream, group, "0").Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("r", r)
	//w, err1 := db.Rdb.XGroupDestroy(ctx, stream, group).Result()
	//fmt.Println("w", w)
	//if err1 != nil {
	//	log.Println(err)
	//	return
	//}
	//result, err := db.Rdb.XInfoGroups(ctx, stream).Result()
	//for i, infoGroup := range result {
	//	fmt.Println(i, infoGroup.Name)
	//}
	//fmt.Println(result)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	// 向流中添加一些数据
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		fields := map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		}
		atoi := strconv.Itoa(i + 330)

		_, err := db.Rdb.XAdd(ctx, &redis.XAddArgs{
			ID:     atoi,
			Stream: stream,
			Values: fields,
		}).Result()
		if err != nil {
			log.Println(err)
			return
		}
	}

	// 消费流中的数据
	for {
		//从流中获取数据
		messages, err := db.Rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumerName,
			Streams:  []string{stream, ">"},
			Count:    10,
			Block:    time.Second * 2,
		}).Result()
		if err != nil {
			log.Println(err)
			messages, err = db.Rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    group,
				Consumer: consumerName,
				Streams:  []string{stream, "0"},
				Count:    10,
				Block:    time.Second * 2,
			}).Result()
		}
		fmt.Println(messages[0])
		fmt.Println(messages[0].Messages)
		/*{my-stream [{2-0 map[field1:value1 field2:value2]} {3-0 map[field1:value1 field2:value2]} {4-0 map[field1:value1 field2:value2]} {5-0 map[field1:value1 field2:value2]} {6-0 map[field1:value1 field2:value2]} {7-0 map[field1:value1 field2:value2]} {8-0 map[field1:value1 field2:value2]} {9-0 map[field1:value1 field2:value2]} {10-0 map[field1:value1 field2:value2]} {21-0 map[field1:value1 field2:value2]}]}
		 */
		//处理数据
		fmt.Println("len(messages[0].Messages):", len(messages[0].Messages))
		list := make([]redis.XMessage, 0)
		fmt.Println(list)
		for i, message := range messages[0].Messages {
			fmt.Println("第", i, "条，ID:", message.ID)
			fmt.Println("第", i, "条，Fields:", message.Values)
			a, err := db.Rdb.XAck(ctx, stream, group, message.ID).Result()
			fmt.Println("a:", a)
			if err != nil {
				panic(err)
			}
			list = append(list, message)
		}
		fmt.Println("list", list)
		fmt.Println("len(messages[0].Messages):", len(messages[0].Messages))
		//fmt.Println(li)
		// 确认已处理的消息

	}
}
