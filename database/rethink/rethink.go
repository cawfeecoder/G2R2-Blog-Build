package rethink

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
)

var (
	session *r.Session
)

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println("Successfully connected to the RethinkDB")
	}
}

//GetSession - Gets Current RethinkDB Session
func GetSession() *r.Session {
	return session
}
