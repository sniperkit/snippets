// Copyright 2015 Jerry Jacobs. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package secdl

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Expire time type
type Expire int64

// Valid URL expire times
const (
	ExpireUnknown   Expire = -1                // Expire time is unknown
	expireSecond           = 1                 // Expire after one second
	ExpireMinute           = 60                // Expire after one minute
	Expire10Minutes        = 10 * ExpireMinute // Expire after 10 minutes
	ExpireHour             = 60 * ExpireMinute // Expire after one hour
	ExpireDay              = 24 * ExpireHour   // Expire after one day
	Expire1Week            = 7 * ExpireDay     // Expire after one week
	Expire2Weeks           = 14 * ExpireDay    // Expire after two weeks
	ExpireMonth            = 30 * ExpireDay    // Expire after one month
	ExpireNever            = 0                 // Expire never
)

// Status code type
type Status int

// Status codes
const (
	StatusValid   Status = 0  // Valid token
	StatusError   Status = -1 // Malformed URL or token
	StatusExpired Status = -2 // Token has expired
)

var expireList = []Expire{Expire10Minutes, ExpireHour, ExpireDay, Expire1Week, Expire2Weeks, ExpireMonth, ExpireNever, expireSecond}

var errExpire = errors.New("invalid Expire")

// SecDl context
type SecDl struct {
	Secret   string
	Filename string
	Prefix   string
	Time     int64
	URL      string
	Expires  Expire
	Status   Status
	BaseURL  string
}

func (e Expire) String() (string, error) {
	var s string

	switch e {
	case ExpireNever:
		s = "never"
	case Expire10Minutes:
		s = "10 minutes"
	case ExpireHour:
		s = "hour"
	case ExpireDay:
		s = "day"
	case Expire1Week:
		s = "1 week"
	case Expire2Weeks:
		s = "2 weeks"
	case ExpireMonth:
		s = "month"
	}

	if s == "" {
		return s, errExpire
	}

	return s, nil
}

// ParseExpire converts a string to a valid expire timespawn
// "n" or "never"        : Expire never
// "10m" or "10 minutes" : Expire after 10 minutes
// "h" or "hour"         : Expire after an hour
// "d" or "day"          : Expire after a day
// "1w" or "1 week"      : Expire after a week
// "2w" or "2 weeks"     : Expire after two weeks
// "m" or "month"        : Expire after a month
func ParseExpire(e string) (Expire, error) {
	v := ExpireUnknown

	switch e {
	case "n", "never":
		v = ExpireNever
	case "10m", "10 minutes":
		v = Expire10Minutes
	case "h", "hour":
		v = ExpireHour
	case "d", "day":
		v = ExpireDay
	case "1w", "1 week":
		v = Expire1Week
	case "2w", "2 weeks":
		v = Expire2Weeks
	case "m", "month":
		v = ExpireMonth
	}

	if v == ExpireUnknown {
		return v, errExpire
	}

	return v, nil
}

// generateToken creates a token from the secret, filename and UNIX timestamp
func generateToken(secret string, filename string, time int64) string {
	p := fmt.Sprintf("%s%s%08x", secret, filename, time)
	hash := md5.Sum([]byte(p))
	return hex.EncodeToString(hash[:])
}

// isExpired checks if the timespawn is expired
func isExpired(t int64, after Expire) bool {
	if after == ExpireNever {
		return false
	} else if time.Now().Unix() < t+int64(after) {
		return false
	}
	return true
}

func expireContains(s []Expire, e Expire) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}
	return false
}

// New creates a SecDl context
func New() *SecDl {
	s := &SecDl{}
	return s
}

// SetBaseURL sets the site base url. E.g "https://example.com"
func (s *SecDl) SetBaseURL(baseurl string) {
	s.BaseURL = baseurl
}

// SetPrefix sets the url prefix, without trailing "/"
func (s *SecDl) SetPrefix(prefix string) {
	s.Prefix = prefix
}

// SetFilename sets the filename, with leading "/"
func (s *SecDl) SetFilename(filename string) {
	s.Filename = filename
}

// SetSecret set the secret
func (s *SecDl) SetSecret(secret string) {
	s.Secret = secret
}

// Encode generates a URL
func (s *SecDl) Encode(e Expire) (url string, err error) {
	if !expireContains(expireList, e) {
		s.Status = StatusError
		return "", fmt.Errorf("Unknown expire specified %d", int64(e))
	}

	s.Time = time.Now().Unix()
	token := generateToken(s.Secret, s.Filename, s.Time+int64(e))
	s.URL = fmt.Sprintf("%s%s%s/%08x%s", s.BaseURL, s.Prefix, token, s.Time, s.Filename)

	return s.URL, nil
}

// Decode verifies a URL
func (s *SecDl) Decode() error {
	url := strings.TrimPrefix(s.URL, s.Prefix)
	parts := strings.SplitN(url, "/", 3)

	if len(parts) < 3 {
		s.Status = StatusError
		return fmt.Errorf("URL is malformed, unexpected field count")
	}

	token := parts[0]
	s.Time, _ = strconv.ParseInt(parts[1], 16, 64)
	s.Filename = "/" + parts[2]

	// Brute force the token with the expireList and known secret
	for i := 0; i < len(expireList); i++ {
		expire := expireList[i]
		_token := generateToken(s.Secret, s.Filename, s.Time+int64(expire))
		if _token == token {
			if isExpired(s.Time, expire) {
				s.Expires = expire
				s.Status = StatusExpired
			} else {
				s.Expires = expire
				s.Status = StatusValid
			}
			return nil
		}
	}

	s.Status = StatusError
	return fmt.Errorf("URL is malformed")
}

// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
func FileServer(secret, prefix, root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpStatus := http.StatusForbidden
		s := New()

		s.SetSecret(secret)
		s.SetPrefix(prefix)
		s.URL = r.URL.Path
		s.Decode()

		if s.Status == StatusValid {
			path := root + s.Filename
			if info, err := os.Stat(path); err == nil && !info.IsDir() {
				http.ServeFile(w, r, path)
				return
			}
		} else if s.Status == StatusExpired {
			httpStatus = http.StatusGone
		}

		http.Error(w, http.StatusText(httpStatus), httpStatus)
	})
}
