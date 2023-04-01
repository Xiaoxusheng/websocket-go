package middleware

import (
	"Gin/db"
	"Gin/models"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

// 消息队列
var stream = "message"

// 生产者
func Producer(message *models.Message) error {
	//binary, err := message.MarshalBinary()
	//if err != nil {
	//	return err
	//}
	w, err := db.Rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		MaxLen: 1000,
		Values: map[string]interface{}{
			message.Room_idently: message,
		},
	}).Result()
	fmt.Println("w", w)
	if err != nil {
		return err
	}
	return nil

}

// 创建消费者群
func CreateConsumerGroup(group string) {
	_, err := db.Rdb.XGroupCreateMkStream(ctx, stream, group, "0").Result()
	if err != nil {
		log.Println("创建消费者群", err)
		return
	}
	//读取
}

// 删除消费者组
func DelConsumerGroup(group string) {
	ctx := context.Background()
	_, err := db.Rdb.XGroupDestroy(ctx, stream, group).Result()
	if err != nil {
		log.Println(err)
		return
	}
}

// 消费者
func Consumer(group string) []redis.XMessage {
	//读取10条消息
	for {
		//"0"表示为确认，">"表示为消费
		result, err := db.Rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:   group,
			Streams: []string{"message", ">"},
			Count:   10,
			Block:   time.Second * 2,
		}).Result()
		if err != nil {
			log.Println(err)
			result, err = db.Rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:   group,
				Streams: []string{"message", "0"},
				Count:   10,
				Block:   time.Second * 2,
			}).Result()
		}
		list := make([]redis.XMessage, 0)
		for _, message := range result[0].Messages {
			fmt.Println("ID:", message.ID)
			fmt.Println("Fields:", message.Values)
			////确认消息
			//_, err := db.Rdb.XAck(ctx, stream, group, message.ID).Result()
			//if err != nil {
			//	log.Println(err)
			//	continue
			//}
			list = append(list, message)
		}
		return list
		//if len(result[0].Messages) > 0 {
		//	_, err := db.Rdb.XAck(ctx, stream, group, result[0].Messages[0].ID).Result()
		//	if err != nil {
		//		log.Println(err)
		//	}
		//}
	}
}

// 确认消息
func Confirmationmessage(group string, id string) {
	_, err := db.Rdb.XAck(ctx, stream, group, id).Result()
	if err != nil {
		log.Println(err)
		return
	}
}
