package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

type Server struct {
	agents    []Agent
	pipelines []*Pipeline
	stages    []*Stage
	jobs      []*Job
	hostname  string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Hostname() string {
	if s.hostname == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return "unset"
		}
		return hostname
	}
	return s.hostname
}

func (s *Server) SetHostname(hostname string) {
	s.hostname = hostname
}

func (s *Server) NewAgent(t AgentType, name string) Agent {
	a := NewAgent(t, name)
	s.AddAgent(a)
	return a
}

func (s *Server) AddAgent(a Agent) {
	s.agents = append(s.agents, a)
}

func (s *Server) AddPipeline(p *Pipeline) {
	s.pipelines = append(s.pipelines, p)
}

func (s *Server) NewPipeline(name string) *Pipeline {
	p := NewPipeline(name)
	s.AddPipeline(p)
	return p
}

func (s *Server) AddJob(j *Job) {
	s.jobs = append(s.jobs, j)
}

func (s *Server) NewJob(name string) *Job {
	j := NewJob(name)
	s.AddJob(j)
	return j
}

func (s *Server) Run() {
	var wg sync.WaitGroup

	for _, pipeline := range s.pipelines {
		wg.Add(1)
		go func(p *Pipeline) {
			defer wg.Done()
			p.Run()
		}(pipeline)
	}

	for _, stage := range s.stages {
		stage.Run()
	}

	for _, job := range s.jobs {
		job.Run()
	}

	wg.Wait()
}

func (s *Server) Summary() {
	fmt.Println("server", s.Hostname(), "summary")
	fmt.Println("---")
	for _, agent := range s.agents {
		ai := agent.Stat()
		fmt.Printf("agent \"%s\" (type %s)\n", agent.Name(), ai.Type())
		fmt.Printf("\tos %s\n", ai.OS())
		fmt.Printf("\tworkspace %s\n", ai.Workspace())
		fmt.Printf("\tfreespace %s\n", humanize.Bytes(ai.Freespace()))
		resources := ai.Resources()
		if len(resources) > 0 {
			fmt.Printf("\tresources\n")
			for _, res := range resources {
				fmt.Println("\t\t-", res)
			}
		}
	}
	fmt.Println("---")
	for _, pipeline := range s.pipelines {
		uuid := s.NewUUID()
		fmt.Printf("pipeline[%s] \"%s\"\n", uuid, pipeline.Name())
	}
	fmt.Println("---")
	for _, stage := range s.stages {
		uuid := s.NewUUID()
		fmt.Printf("stage[%s] \"%s\"\n", uuid, stage.Name())
	}
	fmt.Println("---")
	for _, job := range s.jobs {
		uuid := s.NewUUID()
		fmt.Printf("job[%s] \"%s\"\n", uuid, job.Name())
	}
	fmt.Println("---")
}

func (s *Server) NewUUID() string {
	return uuid.New().String()
}
