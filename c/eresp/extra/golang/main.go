package main

import (
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
)

func main() {
	r, err := redis.Dial("tcp", "192.168.1.12:6379")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer r.Close()

	rep, err := r.Do("PING", true, false, true)
	if err != nil {
		return
	}
	fmt.Printf("%+v",rep)
}
