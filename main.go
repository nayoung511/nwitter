package main

import (
	"os"

	"github.com/nayoung511/blockchain_build/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}