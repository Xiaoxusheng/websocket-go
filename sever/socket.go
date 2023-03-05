package server

import (
	"Gin/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Socket(c *gin.Context) {
	form := c.PostForm("data")
	fmt.Println(form)
	c.JSON(200, gin.H{
		"msg": "成功",
	})

	ctx := context.Background()
	err := db.Rdb.Set(ctx, "name", "leilong", 0).Err()
	if err != nil {
		log.Println(err)
	}
	val := db.Rdb.Get(ctx, "name")
	fmt.Println(val)

}
