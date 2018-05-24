package main

import (
	"log"
	"sync"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/websocket"
)

type wsConn struct {
	con *websocket.Conn
	mu sync.Mutex
	wg sync.WaitGroup
	bgErr error
	pchan chan []byte
	rchan chan []byte
	wchan chan []byte
	cchan chan error
}

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsConnNew(c *websocket.Conn) *wsConn {
	wsc := &wsConn{con: c}
	wsc.pchan = make(chan []byte)
	wsc.rchan = make(chan []byte)
	wsc.wchan = make(chan []byte)
	wsc.cchan = make(chan error)
	return wsc
}

func (c *wsConn) Publish(msg []byte) (err error) {
	c.lock()
	if c.bgErr != nil {
		err = c.bgErr
		c.unlock()
		return err
	}
	c.pchan<-msg
	c.unlock()
	return nil
}

func (c *wsConn) IsClosed() bool {
	_, open := <-c.cchan
	return !open
}

func (c *wsConn) lock() {
	c.mu.Lock()
}

func (c *wsConn) unlock() {
	c.mu.Unlock()
}

func (c *wsConn) Serve() {
	// Read worker
	c.wg.Add(1)
	go func () {
		for {
			_, msg, err := c.con.ReadMessage()
			if err != nil {
				c.cchan<-err
				break
			}
			c.rchan<-msg
		}
		c.wg.Done()
	}()

	// Write worker
	c.wg.Add(1)
	go func() {
		for {
			msg := <-c.wchan
			err := c.con.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				c.cchan<-err
				break
			}
		}
		c.wg.Done()
	}()

	// Superloop
	for {
		select {
		case msg := <-c.rchan:
			log.Println("recv", string(msg))
			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				log.Println("processing...")
				resp, err := http.Post("http://127.0.0.1:1337/ep", "text/json", nil)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				rep, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return
				}
				log.Println("response...")
				c.wchan<-rep
			}()
		case msg := <-c.pchan:
			c.wchan<-msg
		case err := <-c.cchan:
			c.lock()
			c.bgErr = err
			c.unlock()
			log.Println("wsHandlerErr:", err)
			break
		}
	}
}

func (c *wsConn) Close() {
	c.lock()
	c.con.Close()
	c.cchan<-nil
	c.wg.Wait()
	close(c.pchan)
	close(c.rchan)
	close(c.wchan)
	close(c.cchan)
	c.unlock()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}

	ws := wsConnNew(c)
	conchan <- ws
	ws.Serve()
	ws.Close()
}
