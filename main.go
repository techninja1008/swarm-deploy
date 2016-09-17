package main

import (
    "github.com/techninja1008/swarm-deploy/cmd"
    "os"
    "fmt"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}