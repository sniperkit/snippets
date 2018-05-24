package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type LocalAgent struct {
	name      string
	workspace string
	resources *Resources
}

type LocalAgentInfo struct {
	a *LocalAgent
}

func NewLocalAgent(name string) *LocalAgent {
	return &LocalAgent{
		name:      name,
		resources: NewResources(),
	}
}

func (a *LocalAgent) Name() string {
	return a.name
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func diskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

func (a *LocalAgent) AddResource(r *Resource) {
	a.resources.Add(r)
}

func (a *LocalAgent) NewResource(name string) *Resource {
	return a.resources.NewResource(name)
}

func (a *LocalAgent) Workspace() string {
	if a.workspace == "" {
		return os.TempDir()
	}
	return a.workspace
}

func (a *LocalAgent) RunTask(t *Task) error {
	cmd := exec.Command(t.CommandName(), t.CommandArg()...)

	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	return err
}

func (a *LocalAgent) Stat() AgentInfo {
	return &LocalAgentInfo{a: a}
}

func (ai *LocalAgentInfo) Type() AgentType {
	return AgentTypeLocal
}

func (ai *LocalAgentInfo) OS() string {
	return AgentOSUnknown
}

func (ai *LocalAgentInfo) Freespace() uint64 {
	return 0
}

func (ai *LocalAgentInfo) LastSeen() time.Time {
	return time.Time{}
}

func (ai *LocalAgentInfo) Workspace() string {
	return ""
}

func (ai *LocalAgentInfo) Resources() []string {
	return []string{}
}

var _ Agent = &LocalAgent{}
var _ AgentInfo = &LocalAgentInfo{}
