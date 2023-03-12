package test

import (
<<<<<<< HEAD
	"Gin/db"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"os"
	"strconv"
=======
	"fmt"
	"github.com/golang-jwt/jwt/v5"
>>>>>>> origin/master
	"testing"
	"time"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func Test_createtoken(*testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user{"123", "448", jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), Issuer: "test"}})
	signedString, err := token.SignedString([]byte("suhusduhhs65+978t$67"))
	if err != nil {
		return
	}
	fmt.Println(signedString)
}

func Test_parseToken(*testing.T) {
<<<<<<< HEAD
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMyIsInBhc3N3b3JkIjoiNDQ4IiwiaXNzIjoidGVzdCIsImV4cCI6MTY3ODQ1NzY2MX0.rEpVxcaOcqOwLh8wKPvuzTHAGdIBTv1OaACezqlcpCo"
=======
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMyIsInBhc3N3b3JkIjoiNDQ4IiwiaXNzIjoidGVzdCIsImV4cCI6MTY3ODA5ODQ0OH0.n5llexjpe1rzHoph0XeN4BZIQraLUo88BAEQTq2ox1s"
>>>>>>> origin/master
	token, err := jwt.ParseWithClaims(tokenString, &user{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("suhusduhhs65+978t$67"), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*user); ok && token.Valid {
		fmt.Printf("%v %v %v", claims.Username, claims.Password, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}
}
<<<<<<< HEAD

func Test_1(*testing.T) {
	err := os.Setenv("key", ("suhusduhhs65+978t$67"))
	if err != nil {
		return
	}
	fmt.Println(os.Getenv("key"))
	fmt.Println(time.Now().Unix())
}

func Test_2(*testing.T) {
	ctx := context.Background()

	////五分钟后过期
	////db.Rdb.Set(ctx, "randString", utility.Emails(), time.Second*60*5)
	//
	//a := map[string]interface{}{
	//	"randString": 1, "time": time.Now().Unix(), "username": "123"}
	//
	//result, err := db.Rdb.HSet(ctx, "email", a).Result()
	//db.Rdb.Expire(ctx, "email", time.Second*10)
	////
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	////fmt.Println(result)
	////fmt.Println(db.Rdb.HMGet(ctx, "email", "username", "time", "randString").Result())
	//results, errs := db.Rdb.Exists(ctx, "email").Result()
	//if errs != nil {
	//	return
	//}
	//fmt.Println(results)
	//if result == 0 {
	//	return
	//}
	v, _ := db.Rdb.HMGet(ctx, "email", "username", "time", "randString").Result()
	hset := make([]string, 0)
	fmt.Println(v)

	for i, i2 := range v {
		fmt.Println(i, i2.(string))
		hset = append(hset, i2.(string))
	}
	fmt.Println(hset)
	db.Rdb.Del(ctx, "lei")

}

func Test_hehe(*testing.T) {
	rand.Seed(time.Now().UnixMicro())
	//七位数
	rand.Int63n(10)
	str := ""
	for i := 0; i < 7; i++ {
		str += strconv.FormatInt(rand.Int63n(10), 10)
	}
	fmt.Println(time.Now().Unix())
}

/*func OperateDB() {
	// 查看所有的键
	res, err := DB.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// 查看某一个键是否存在
	n, err := DB.Exists(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	// 查看键的类型
	str, err := DB.Type(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	// 为键设置过期时间
	success, err := DB.Expire(ctx, "name", 5*time.Minute).Result()
	if err != nil {
		fmt.Println(err)
	}
	if success {
		fmt.Println("设置时间成功")
	}
	//查看键的过期时间
	t, err := DB.TTL(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	time.Sleep(5 * time.Second)
	t, err = DB.TTL(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	//更换数据库
	n, err = DB.DBSize(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库的大小为", n)

	//删除键
	n, err = DB.Del(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("删除成功")

	//删除库中所有的数据
	DB.FlushDB(ctx).Result()

}*/
=======
>>>>>>> origin/master
