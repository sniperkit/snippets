package redislog

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisLog writes loglines to Redis
type RedisLog struct {
	c   redis.Conn
	mu  sync.Mutex
	key string
}

// New creates a new logger which connects to URI and writes to key
func New(URI, key string) (*RedisLog, error) {
	c, err := redis.Dial("tcp", URI)
	if err != nil {
		return nil, err
	}

	rl := &RedisLog{
		c:   c,
		key: key,
	}
	return rl, nil
}

// Close the logger redis connection
func (rl *RedisLog) Close() error {
	return rl.c.Close()
}

// log message
func (rl *RedisLog) log(msg string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// nolint
	if !strings.HasSuffix(msg, "\n") {
		_, _ = rl.c.Do("RPUSH", rl.key, msg+"\n")
	} else {
		_, _ = rl.c.Do("RPUSH", rl.key, msg)
	}
}

// Stderr allocates a io.Writer for loglevel error
func (rl *RedisLog) Stderr() io.Writer {
	rp, wp := io.Pipe()
	go func() {
		defer rp.Close()
		scanner := bufio.NewScanner(rp)
		for scanner.Scan() {
			rl.Error(scanner.Text())
		}
	}()
	return wp
}

// Stdout allocates a io.Writer for loglevel info
func (rl *RedisLog) Stdout() io.Writer {
	rp, wp := io.Pipe()
	go func() {
		defer rp.Close()
		scanner := bufio.NewScanner(rp)
		for scanner.Scan() {
			rl.Info(scanner.Text())
		}
	}()
	return wp
}

// Debug writes log message in level debug
func (rl *RedisLog) Debug(args ...interface{}) {
	rl.log(fmt.Sprint(args...))
}

// Debugf writes a formatted log message in level debug
func (rl *RedisLog) Debugf(format string, args ...interface{}) {
	rl.log(fmt.Sprintf(format, args...))
}

// Debugln writes a log message in level debug with newline
func (rl *RedisLog) Debugln(args ...interface{}) {
	rl.log(fmt.Sprintln(args...))
}

// Info writes log message in level info
func (rl *RedisLog) Info(args ...interface{}) {
	rl.log(fmt.Sprint(args...))
}

// Infof writes a formatted log message in level info
func (rl *RedisLog) Infof(format string, args ...interface{}) {
	rl.log(fmt.Sprintf(format, args...))
}

// Infoln writes a log message in level info with newline
func (rl *RedisLog) Infoln(args ...interface{}) {
	rl.log(fmt.Sprintln(args...))
}

// Warning writes log message in level warning
func (rl *RedisLog) Warning(args ...interface{}) {
	rl.log(fmt.Sprint(args...))
}

// Warningf writes a formatted log message in level warning
func (rl *RedisLog) Warningf(format string, args ...interface{}) {
	rl.log(fmt.Sprintf(format, args...))
}

// Warningln writes log message in level warning with newline
func (rl *RedisLog) Warningln(args ...interface{}) {
	rl.log(fmt.Sprintln(args...))
}

// Error writes log message in level error
func (rl *RedisLog) Error(args ...interface{}) {
	rl.log(fmt.Sprint(args...))
}

// Errorf writes a formatted log message in level error {
func (rl *RedisLog) Errorf(format string, args ...interface{}) {
	rl.log(fmt.Sprintf(format, args...))
}

// Errorln writes log message in level error with newline
func (rl *RedisLog) Errorln(args ...interface{}) {
	rl.log(fmt.Sprintln(args...))
}

// Subscribe for log messages and write to msg channel
func (rl *RedisLog) Subscribe(ctx context.Context, wc io.WriteCloser) {
	go func() {
		defer wc.Close()

		llen := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				nllen, err := rl.pullLoglines(wc, llen)
				if err != nil {
					return
				}
				llen = nllen
			}
		}
	}()
}

func (rl *RedisLog) pullLoglines(w io.Writer, llen int) (int, error) {
	// Poll for new entries
	rl.mu.Lock()
	nllen, err := redis.Int(rl.c.Do("LLEN", rl.key))
	rl.mu.Unlock()
	if err != nil {
		return llen, err
	}

	if nllen <= llen {
		return llen, nil
	}

	// Get new loglines
	rl.mu.Lock()
	nll, err := redis.ByteSlices(rl.c.Do("LRANGE", rl.key, llen, nllen))
	rl.mu.Unlock()

	if err != nil {
		return llen, err
	}

	// Write loglines
	for _, ll := range nll {
		_, err = w.Write(ll)
		if err != nil {
			return 0, err
		}
	}

	return nllen, nil
}
