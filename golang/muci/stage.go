package main

import (
	"fmt"
	"time"
)

type Stage struct {
	name      string
	jobs      []*Job
	time      time.Time
	workspace string
}

func NewStage(name string) *Stage {
	return &Stage{name: name}
}

func (s *Stage) AddJob(j *Job) {
	s.jobs = append(s.jobs, j)
}

func (s *Stage) SetWorkspace(workspace string) {
	s.workspace = workspace
}

func (s *Stage) Workspace() string {
	return s.workspace
}

func (s *Stage) NewJob(name string) *Job {
	j := NewJob(name)
	s.AddJob(j)
	return j
}

func (s *Stage) SetName(name string) {
	s.name = name
}

func (s *Stage) Time() time.Time {
	return s.time
}

func (s *Stage) Name() string {
	return s.name
}

func (s *Stage) Run() error {
	s.time = time.Now()

	if len(s.jobs) == 0 {
		fmt.Printf("stage[%s] no jobs to run", s.name)
		return nil
	}

	for _, job := range s.jobs {
		fmt.Printf("stage[%s]: job \"%s\" started\n", s.name, job.Name())
		job.Run()
		fmt.Printf("stage[%s]: job \"%s\" finished\n", s.name, job.Name())
	}

	return nil
}

var _ Runner = &Stage{}
