package cmd

import (
    "github.com/spf13/cobra"
    "github.com/docker/docker/cli/command/stack"
    "github.com/techninja1008/swarm-deploy/swarm"
)

var (
    RootCmd *cobra.Command = &cobra.Command{
        Use: "swarm-deploy",
        Short: "swarm-deploy is a tool for managing docker swarm-mode clusters",
    }
    GoCmd *cobra.Command = &cobra.Command{
        Use: "go",
        Short: "fully initialize the project in the current folder",
        Run: func(cmd *cobra.Command, args []string){
            
        },
    }
)

func init() {
    RootCmd.AddCommand(
        GoCmd,
        stack.NewTopLevelDeployCommand(swarm.NewDockerCli())
    )
}
