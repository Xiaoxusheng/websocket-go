package models

import "Gin/db"

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

// 查询每小时访问最频繁ip
func GetIPNumber(time int64) (*[]Ip, error) {
	iplist := []Ip{}
	err := db.DB.Select(&iplist, "select *from ip where time>? ", time)
	if err != nil {
		return nil, err
	}
	return &iplist, nil
}
