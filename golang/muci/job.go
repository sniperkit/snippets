package main

import (
	"fmt"
	"time"
)

type Job struct {
	name      string
	agent     Agent
	tasks     []*Task
	time      time.Time
	resources *Resources
}

func NewJob(name string) *Job {
	return &Job{
		name:      name,
		resources: NewResources(),
	}
}

func (j *Job) AddTask(t *Task) {
	j.tasks = append(j.tasks, t)
}

func (j *Job) NewTask(name string) *Task {
	t := NewTask(name)
	j.AddTask(t)
	return t
}

func (j *Job) SetAgent(a Agent) {
	j.agent = a
}

func (j *Job) AddResource(r *Resource) {
	j.resources.Add(r)
}

func (j *Job) NewResource(name string) *Resource {
	r := j.resources.NewResource(name)
	j.AddResource(r)
	return r
}

func (j *Job) Name() string {
	return j.name
}

func (j *Job) Time() time.Time {
	return j.time
}

func (j *Job) Run() error {
	j.time = time.Now()

	if len(j.tasks) == 0 {
		fmt.Printf("job[%s]: no tasks to run\n", j.name)
		return nil
	}

	if j.agent == nil {
		fmt.Printf("job[%s]: no agent assigned\n", j.name)
		return nil
	}

	fmt.Printf("job[%s]: running on agent \"%s\"\n", j.name, j.agent.Name())

	for _, task := range j.tasks {
		fmt.Printf("job[%s]: task \"%s\" started\n", j.name, task.Name())
		fmt.Printf("         command: \"%s\"\n", task.CommandString())
		err := task.Run(j.agent)
		if err != nil {
			fmt.Printf("job[%s]: task \"%s\" error: %v\n", j.name, task.Name(), err)
		}
		fmt.Printf("         task took %v\n", time.Since(task.Time()))
		fmt.Printf("job[%s]: task \"%s\" finished\n", j.name, task.Name())
		if err != nil {
			break
		}
	}
	fmt.Printf("job[%s] took %v\n", j.name, time.Since(j.time))

	return nil
}

var _ Runner = &Job{}
