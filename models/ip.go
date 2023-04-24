package models

import (
	"Gin/db"
	"time"
)

type Ip struct {
	Id          int    `json:"id"`
	Ip          string `json:"ip"`
	Time        int64  `json:"time"`
	Useindently string `json:"useindently"`
}

type IPs struct {
	Ip          string `json:"ip"`
	Time        int64  `json:"time"`
	Useindently string `json:"useindently"`
}

func InsertIpbyUser(ip *IPs) error {
	_, err := db.DB.Exec("insert into ip(ip,time,useindently) value (?,?,?)", ip.Ip, ip.Time, ip.Useindently)
	if err != nil {
		return err
	}
	return nil
}

// 查询ip
func GetIPNumber(ip string) (int, error) {
	type numer struct {
		Num int `json:"num"`
	}

	iplist := numer{}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()) // 今天0点的时间
	timestamp := today.Unix()                                                                                      // 转换为时间戳
	err := db.DB.Get(&iplist, "select count(ip) as num from ip  where time > ? and ip=?", timestamp, ip)
	if err != nil {
		return 0, err
	}
	return iplist.Num, nil
}
