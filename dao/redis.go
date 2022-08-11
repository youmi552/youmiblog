package dao

import "github.com/gomodule/redigo/redis"

var rd redis.Conn

func init() {
	//conn, err := redis.Dial("tcp", ":6379")
	//if err != nil {
	//	panic(err)
	//}
	//rd = conn
}
