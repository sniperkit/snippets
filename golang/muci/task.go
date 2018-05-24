package main

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	name    string
	time    time.Time
	cmdName string
	cmdArg  []string
}

func NewTask(name string) *Task {
	return &Task{name: name}
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) SetCommand(name string, arg ...string) {
	t.cmdName = name
	t.cmdArg = arg
}

func (t *Task) CommandName() string {
	return t.cmdName
}

func (t *Task) CommandArg() []string {
	return t.cmdArg
}

func (t *Task) CommandString() string {
	return fmt.Sprintf("%s %s", t.cmdName, strings.Join(t.cmdArg, " "))
}

func (t *Task) Time() time.Time {
	return t.time
}

func (t *Task) Run(a Agent) error {
	t.time = time.Now()
	return a.RunTask(t)
}
