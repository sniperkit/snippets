package main

import (
	"time"
	"net/http"
)

func helloworld(c *wsConn) {
	var cnt uint
	for {
		if err := c.Publish([]byte("boem")); err != nil {
			return
		}
		time.Sleep(time.Second)
		cnt++
		if cnt == 3 {
			if err := c.Publish([]byte("bye")); err != nil {
				return
			}
			return
		}
	}
}

var conchan chan *wsConn

func main() {
	http.HandleFunc("/ep", epHandler)
	http.HandleFunc("/ws", wsHandler)

	conchan = make(chan *wsConn)

	go func() {
		c := <-conchan
		go func() {
			helloworld(c)
		}()
	}()

	http.ListenAndServe("0.0.0.0:1337", nil)
}
