package dab

import (
    "github.com/docker/docker/cli/command"
    "github.com/docker/docker/cli/command/stack"
    "github.com/techninja1008/swarm-deploy/swarm"
)

func DeployC(dockerCli *command.DockerCli, filename string, ns string) error {
    deployCmd := stack.NewTopLevelDeployCommand(dockerCli)
    
    deployCmd.SetArgs([]string{"--with-registry-auth", "--file", filename, ns})
    
    return deployCmd.Execute()
}

func Deploy(filename string, ns string) error {
    return DeployC(swarm.NewDockerCli(), filename, ns)
}