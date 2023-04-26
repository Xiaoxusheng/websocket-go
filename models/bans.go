package models

import (
	"Gin/db"
)

type Bans struct {
	IP   string
	Time string
}
type SearchIP struct {
	Id   int
	IP   string
	Time string
}

// 封杀ip
func BanIP(ip *Bans) error {
	_, err := db.DB.Exec("insert into bans(ip,time) value (?,?)", ip.IP, ip.Time)
	if err != nil {
		return err
	}
	return nil
}

// 查询是否为封杀ip
func GetbanIp(ip string) bool {
	bans := SearchIP{}
	err := db.DB.Get(&bans, "select * from bans where ip=?", ip)
	if err != nil {
		return false
	}
	return true
}
