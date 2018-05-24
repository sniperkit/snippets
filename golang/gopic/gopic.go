// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/xor-gate/gopic/lib/config"
	"github.com/xor-gate/gopic/lib/db"
	"time"
)

func main() {
	cfg := config.Get()

	err := db.Open(cfg.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	httpHandler()

	fmt.Println("Running gopic", config.Version())
	fmt.Println(fmt.Sprintf("Serving at http://%s:%s", cfg.Host, cfg.Port))

	var stats IndexerStat

	// Indexing
	for {
		fmt.Println("Indexing pictures...this will take some time")

		stats.Start()
		Indexer(db.Get(), cfg.Path, &stats)
		stats.Finish()

		fmt.Println(stats)
		stats.Reset()

		fmt.Println("\nSleeping for",cfg.PollTime,"minutes...")

		time.Sleep(time.Minute * time.Duration(cfg.PollTime))
	}
}
