package main

import (
	"github.com/AndrusGerman/workspace-runner/cmd/workspace-runner/bootstrap"
	"github.com/AndrusGerman/workspace-runner/internal/adapters/server"
)

func mainServer(bootstrap *bootstrap.Bootstrap) {
	var server = server.NewServer(bootstrap.WorkspaceService, bootstrap.ProjectService)
	server.Start()
}
