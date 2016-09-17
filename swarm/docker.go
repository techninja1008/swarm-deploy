package swarm

import (
    "github.com/docker/docker/pkg/term"
    "github.com/docker/docker/cli/command"
)

func NewDockerCli() *command.DockerCli {
    stdin, stdout, stderr := term.StdStreams()
    return command.NewDockerCli(stdin, stdout, stderr)
}