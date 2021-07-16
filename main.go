package main

import (
	"gochain/cli"
	"os"
)



func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()

}