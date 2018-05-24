package main

import (
	"fmt"
)

func main() {
	srv := NewServer()

	// Add ssh agent
	sshagent := srv.NewAgent(AgentTypeSSH, "muci@192.168.1.201:22").(*SSHAgent)
	sshagent.NewResource("debian")
	sshagent.NewResource("linux")
	sshagent.SetUser("muci")
	sshagent.SetWorkspace("~/workspace")
	sshagent.SetHost("192.168.1.201:22")
	sshagent.SetPrivateKeyFromFile("./id_rsa")

	// Create new pipeline
	pipeline1 := srv.NewPipeline("pipeline1")

	// Create new stage1 in pipeline1
	stage1 := pipeline1.NewStage("stage1")
	{
		workspace := sshagent.Workspace() + "/" + srv.NewUUID()
		fmt.Println("workspace:", workspace)

		// Add job1 to stage1
		job1 := stage1.NewJob("stage1-job1")
		// TODO manual assigned agent, on job run it should be automaticly assigned
		//      the server is aware of the available agents with resources
		job1.SetAgent(sshagent)

		// Create workspace
		job1.NewTask("stage1-job1-task1").SetCommand("/bin/mkdir", "-p", workspace)

		// Clone github.com/xor-gate/muci

		// Remove workspace
		job1.NewTask("stage1-job1-task1").SetCommand("/bin/rm", "-rf", workspace)
	}

	// Print server summary
	srv.Summary()

	// Run all server registered pipelines, stages, tasks
	srv.Run()
}
