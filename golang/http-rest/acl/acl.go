package acl

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
	"errors"
	"gopkg.in/redis.v3"
)

type Token struct {
	Id string       `json:"id"`
	Ctime time.Time `json:"ctime,omitempty"`
	Atime time.Time `json:"atime,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}
type Tokens []Token

type Acl struct {
	store *AclStore
	user string  // User X-Xg-Auth-User
	ReqToken *Token `json:"-"` // Token X-Xg-Api-Token
	TokenList Tokens `json:"tokens"`
	Permissions []string `json:"permissions"`
}

type AclStore struct {
	db *redis.Client
}

func (t *Token) String() string {
	j, _ := json.Marshal(t)
	return string(j)
}

func (t *Token) Bytes() []byte {
	return []byte(t.String())
}

func (t Tokens) String() string {
	j, _ := json.Marshal(t)
	return string(j)
}

func (t Tokens) Bytes() []byte {
	return []byte(t.String())
}

func (a *Acl) GetToken(id string) (*Token, int) {
	for i := range a.TokenList {
		if a.TokenList[i].Id == id {
			return &a.TokenList[i],i
		}
	}
	return nil,0
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
func tokenRandId(length int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Create new token and sync with AclStore
func (a *Acl) NewToken() *Token {
	t := &Token{}

	t.Id    = tokenRandId(32)
	t.Ctime = time.Now()

	a.TokenList = append(a.TokenList, *t)
	a.Save()

	return t
}

// Create new token and sync with AclStore
func (a *Acl) DeleteToken(id string) bool {
	t, idx := a.GetToken(id)
	if t == nil {
		fmt.Println("Delete fail", id)
		return false
	}

	a.TokenList[idx] = a.TokenList[len(a.TokenList)-1]
	a.TokenList = a.TokenList[0:len(a.TokenList)-1]

	fmt.Println("Deleted token", id)

	a.Save()

	return true
}

func (a *Acl) Save() {
	a.store.Save(a)
}

func (s *AclStore) Save(a *Acl) {
	j, _ := json.Marshal(a)
	err := s.db.Set(a.user, j, 0).Err()
	if err != nil {
		fmt.Println("AclStoreSave", err)
	}
}

func (s *AclStore) Load(user string) (*Acl, error) {
	// Load data from redis "user" : "<json>"
	k := s.db.Get(user)
	v, err := k.Result()
	if err == redis.Nil {
		return nil,errors.New("omg")
	}

	// Decode JSON
	var a Acl
	err = json.Unmarshal([]byte(v), &a)
	if err != nil {
		return nil,err
	}

	a.user = user

	return &a,nil
}

func (s *AclStore) Auth(user, token, path string) *Acl {
	if user == "" || token == "" || path == "" {
		return nil
	}

	// Load
	a, err := s.Load(user)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Load token
	t, _ := a.GetToken(token)
	if t == nil {
		fmt.Println("token doesnt exist: ", token)
		return nil
	}

	// Verify
	if t.Id != token {
		return nil
	}

	// Update token access time
	t.Atime = time.Now()
	s.Save(a)

	// Attach the store and current used token
	a.store = s
	a.ReqToken = t

	return a
}

func New(host string) (AclStore, error) {
	a := AclStore{}

	a.db = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return a,nil
}
