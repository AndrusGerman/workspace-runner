package main

import (
	"flag"

	"github.com/AndrusGerman/workspace-runner/cmd/workspace-runner/bootstrap"
	"github.com/AndrusGerman/workspace-runner/internal/adapters/server"
	"github.com/AndrusGerman/workspace-runner/internal/core/services"
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
		var templateService, err = services.NewTemplateService()
		if err != nil {
			panic(err)
		}
		var server = server.NewServer(bootstrap.WorkspaceService, templateService, bootstrap.ProjectService)
		server.Start()
		return
	}

	hello()
}
