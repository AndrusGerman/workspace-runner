package main

import (
	"flag"

	"github.com/AndrusGerman/workspace-runner/cmd/workspace-runner/bootstrap"
)

func init() {
	flag.Parse()
}

func main() {
	var bootstrap, err = bootstrap.NewBootstrap()
	if err != nil {
		panic(err)
	}

	if len(flag.Args()) > 0 && flag.Args()[0] == "run" {
		run(bootstrap)
		return
	}

	if len(flag.Args()) > 0 && flag.Args()[0] == "list" {
		list(bootstrap)
		return
	}

	if len(flag.Args()) > 0 && flag.Args()[0] == "server" {
		mainServer(bootstrap)
		return
	}

	hello()
}
