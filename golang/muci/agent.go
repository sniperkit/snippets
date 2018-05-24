package main

import (
	"time"
)

type Agent interface {
	Name() string
	Stat() AgentInfo
	RunTask(t *Task) error
}

type AgentInfo interface {
	Type() AgentType
	OS() string
	Freespace() uint64
	Workspace() string
	LastSeen() time.Time
	Resources() []string
}

type AgentType string

const (
	AgentTypeLocal  = "local"
	AgentTypeSSH    = "ssh"
	AgentTypeRemote = "remote"
)

const AgentOSUnknown = "unknown"

func NewAgent(t AgentType, name string) Agent {
	switch t {
	case AgentTypeLocal:
		return NewLocalAgent(name)
	case AgentTypeSSH:
		return NewSSHAgent(name)
	case AgentTypeRemote:
		return NewRemoteAgent(name)
	}
	return nil
}
