// Copyright 2017 GoPic Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package config

import (
	"flag"
)

// Config from CLI application
type Config struct {
	DB          string
	Path        string
	Host        string
	Port        string
	AdminSecret string
	PollTime    int
}

var (
	//Version contains the application version
	version string
	//Build contains the repository branch and commit hash used to
	//compile the agent.
	build string
	// GoPic configuration
	config Config
)

func init() {
	// Load defaults
	var db   = flag.String("db", "gopic.db", "Filepath to index database")
	var path = flag.String("path", "", "Path to add to index")
	var host = flag.String("host", "0.0.0.0", "Host to bind the HTTP server")
	var port = flag.String("port", "8081", "Port to bind the HTTP server")
	var pollTime = flag.Int("pollTimeMinutes", 15, "Reindex poll time in minutes")

	// Parse cli flags
	flag.Parse()

	// Load defaults or overwritten config
	config.DB   = *db
	config.PollTime = *pollTime
	config.Path = *path
	config.Host = *host
	config.Port = *port
}

// Get a snapshot of the configuration
func Get() Config {
	return config
}

// Version contains the semantic version of the agent build
func Version() string {
	if version != "" {
		return version
	}
	return "dev"
}

// Build is the git branch and commit hash provided during compilation
func Build() string {
	return build
}

// Host returns the configured HTTP host to bind to
func Host() string {
	return config.Host
}

// Port returns the configured HTTP port to bind to
func Port() string {
	return config.Port
}

// DB returns the filepath to the database
func DB() string {
	return config.DB
}
