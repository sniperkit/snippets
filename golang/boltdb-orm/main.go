package main

import (
	"fmt"
	"github.com/asdine/storm"
)

func main() {
	db, err := storm.Open("example.boltdb")
	if err != nil {
		return
	}
	defer db.Close()

	var cfg Config

	cfg.Load(db)

	fmt.Printf("%+v\n", cfg)
	fmt.Printf("port: %d\n", cfg.Port)

	cfg.Port = 9999
	cfg.Save(db)
}
