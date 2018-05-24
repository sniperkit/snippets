package main

import (
	"github.com/asdine/storm"
)

type Config struct {
	Port uint16 `storm:"port"`
}

var cfgDefaults = Config {
	Port: 1337,
}

func (cfg *Config) Load(db *storm.DB) {
	err := db.Get("config", "config", cfg)
	if err != nil {
		db.Set("config", "config", &cfgDefaults)
		*cfg = cfgDefaults
	}
}

func (cfg *Config) Save(db *storm.DB) {
	db.Set("config", "config", cfg)
}
