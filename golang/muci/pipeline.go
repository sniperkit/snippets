package main

import (
	"fmt"
	"sync"
	"time"
)

type Pipeline struct {
	name   string
	stages []*Stage
	time   time.Time
}

func NewPipeline(name string) *Pipeline {
	return &Pipeline{name: name}
}

func (p *Pipeline) Name() string {
	return p.name
}

func (p *Pipeline) AddStage(s *Stage) {
	p.stages = append(p.stages, s)
}

func (p *Pipeline) NewStage(name string) *Stage {
	s := NewStage(name)
	p.AddStage(s)
	return s
}

func (p *Pipeline) Time() time.Time {
	return p.time
}

func (p *Pipeline) Run() error {
	p.time = time.Now()

	if len(p.stages) == 0 {
		fmt.Printf("pipeline[%s]: no stages to run\n", p.name)
		return nil
	}

	var wg sync.WaitGroup

	for _, stage := range p.stages {
		wg.Add(1)
		go func(s *Stage) {
			defer wg.Done()
			fmt.Printf("pipeline[%s]: stage \"%s\" started\n", p.name, s.Name())
			s.Run()
			fmt.Printf("pipeline[%s]: stage \"%s\" finished (took %v)\n", p.name, stage.Name(), time.Since(stage.Time()))
		}(stage)
	}
	fmt.Printf("pipeline[%s] took %v\n", p.name, time.Since(p.time))

	wg.Wait()

	// TODO collect errors
	return nil
}

var _ Runner = &Pipeline{}
