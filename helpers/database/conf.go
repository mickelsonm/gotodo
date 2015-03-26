package database

import (
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

func MongoConnectionString() *mgo.DialInfo {
	info := new(mgo.DialInfo)
	addr := os.Getenv("MONGO_URL")
	if addr == "" {
		addr = "127.0.0.1"
	}

	info.Addrs = append(info.Addrs, addr)
	info.Username = os.Getenv("MONGO_CART_USERNAME")
	info.Password = os.Getenv("MONGO_CART_PASSWORD")
	info.Database = os.Getenv("MONGO_CART_DATABASE")
	info.Timeout = time.Second * 2
	info.FailFast = true
	if info.Database == "" {
		info.Database = "TodosDB"
	}
	return info
}
