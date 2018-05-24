package main

import (
	"time"
)

type RemoteAgent struct {
	name      string
	workspace string
	resources *Resources
}

type RemoteAgentInfo struct {
	a *RemoteAgent
}

func NewRemoteAgent(name string) *RemoteAgent {
	return &RemoteAgent{
		name:      name,
		resources: NewResources(),
	}
}

func (a *RemoteAgent) Name() string {
	return a.name
}

func (a *RemoteAgent) AddResource(r *Resource) {
	a.resources.Add(r)
}

func (a *RemoteAgent) NewResource(name string) *Resource {
	return a.resources.NewResource(name)
}

func (a *RemoteAgent) Workspace() string {
	return a.workspace
}

func (a *RemoteAgent) RunTask(t *Task) error {
	return nil
}

func (a *RemoteAgent) Stat() AgentInfo {
	return &RemoteAgentInfo{a: a}
}

func (ai *RemoteAgentInfo) Type() AgentType {
	return AgentTypeRemote
}

func (ai *RemoteAgentInfo) OS() string {
	return AgentOSUnknown
}

func (ai *RemoteAgentInfo) Freespace() uint64 {
	return 0
}

func (ai *RemoteAgentInfo) LastSeen() time.Time {
	return time.Time{}
}

func (ai *RemoteAgentInfo) Workspace() string {
	return ""
}

func (ai *RemoteAgentInfo) Resources() []string {
	return []string{}
}

var _ Agent = &RemoteAgent{}
var _ AgentInfo = &RemoteAgentInfo{}
