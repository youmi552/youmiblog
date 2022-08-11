package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func checkErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
	}
}
func main() {
	//conn, _ := redis.Dial("tcp", ":6379")
	//defer conn.Close()
	//conn.Do("set", "c1", "hello2")
	conn, err := redis.Dial("tcp", ":6379")
	defer conn.Close()
	checkErr(err)
	//写入数据
	conn.Do("set", "user:name", "youmi")
	//读取数据
	username, err := redis.String(conn.Do("get", "user:name"))
	checkErr(err)
	fmt.Println(username)
	//检查数据是否存在
	b, err := redis.Bool(conn.Do("exists", "user:name"))
	checkErr(err)
	if b {
		fmt.Println("username存在")
	} else {
		fmt.Println("username不存在")
	}
	//设置redis过期时间3s
	_, err = conn.Do("set", "myName", "hehe", "ex", 3)
	checkErr(err)
	myName, err := redis.String(conn.Do("get", "myName"))
	fmt.Println("myName : ", myName)
	//5s后取数据
	time.Sleep(time.Second * 5)
	myName, err = redis.String(conn.Do("get", "myName"))
	if err != nil {
		fmt.Println("After 5s ", err)
	} else {
		fmt.Println("After 5s myName : ", myName)
	}
}
