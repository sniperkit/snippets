package main

import (
	"io"
	"bufio"
	"strings"
	"bytes"

	"github.com/garyburd/redigo/redis"
)

func main() {
	rl, err := New(":6379", "testlog")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	io.Copy(rl, strings.NewReader("hello world!\nhello world...!"))
}

type RedisLogger struct {
	c redis.Conn
	logKey string
}

func (rl *RedisLogger) Write(p []byte) (n int, err error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	for s.Scan() {
		rl.c.Do("LPUSH", rl.logKey, s.Text())
	}
	return len(p), nil
}

func New(URI, logKey string) (*RedisLogger, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}

	rl := &RedisLogger {
		c: c,
		logKey: logKey,
	}
	return rl, nil
}

func (rl *RedisLogger) Close() error {
	return rl.c.Close()
}
