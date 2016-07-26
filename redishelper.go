package redishelper

import (
	"log"
	"github.com/garyburd/redigo/redis"
)

var RedisConn redis.Conn
var err error

func Connect() {
	/*
	 * Redis Connection
	 */
	RedisConn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("Redis connection failed")
	}
}

/**
 * Helper function for redis to run a command and convert the output into []bytes
 */
func ExecCommand(commandName string, args ...interface{}) (reply []byte, err error) {

	/*
	 * Redis Connection
	 */
	data, err := redis.Bytes(RedisConn.Do(commandName, args...))
	if err != nil {
		log.Printf("Redis Error: %#v\n", err)
	}

	return data, err
}