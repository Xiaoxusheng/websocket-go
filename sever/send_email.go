package server

import (
	"Gin/db"
	"Gin/utility"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 发送验证码
func Send_email(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	code := utility.Getcode()
	ctx := context.Background()
	//键值是否存在
	res, err := db.Rdb.Exists(ctx, username).Result()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("res", res)
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "姓名不能为空！",
		})
		return
	}

	//验证码不存在
	if res == 0 {
		reslust, err := db.Rdb.HSet(ctx, username, "randString", code, "time", time.Now().Unix(), "username", username).Result()
		fmt.Println(reslust)
		//1分钟后过期
		db.Rdb.Expire(ctx, username, time.Second*60)
		fmt.Println(db.Rdb.TTL(ctx, username).Result())
		if err != nil {
			log.Println("插入数据出错：" + err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "获取验证码成功！",
		})
		utility.Sendemails()
	} else {
		result, err := db.Rdb.HMGet(ctx, username, "username", "time", "randString").Result()
		fmt.Println(db.Rdb.TTL(ctx, username).Result())
		fmt.Println(result)
		if err != nil {
			return
		}
		hset := make([]string, 0)
		for i, v := range result {
			fmt.Println(i, "v", v)
			hset = append(hset, v.(string))
		}
		log.Println(hset)
		if username == hset[0] {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "请求过于频繁，请稍后在试！",
				"code": 1,
			})
			return
		}

	}

}
