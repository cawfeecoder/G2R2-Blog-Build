package redis

import (
	"fmt"
	"log"

	"gopkg.in/redis.v5"
)

var (
	session *r.Session
)

func init() {
	var err error

	session, err = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })
	pong, err := client.Ping().Result()
  if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println("Successfully connected to the Redis")
	}

  //GetSession - Gets Current RethinkDB Session
  func GetSession() *redis.Client {
	  return session
  }
}
