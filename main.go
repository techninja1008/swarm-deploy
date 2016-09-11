package main

import (
    "github.com/spf13/cobra"
    "github.com/techninja1008/swarm-deploy/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}