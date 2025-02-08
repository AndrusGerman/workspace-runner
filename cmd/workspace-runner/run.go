package main

import (
	"context"
	"flag"
	"log"

	"github.com/AndrusGerman/workspace-runner/cmd/workspace-runner/bootstrap"
)

func run(bootstrap *bootstrap.Bootstrap) {
	if len(flag.Args()) < 2 {
		log.Println("No workspace name provided")
		return

	}
	var workspaceName = flag.Args()[1]

	workspace, err := bootstrap.WorkspaceService.GetByName(context.Background(), workspaceName)

	if err != nil {
		panic(err)
	}

	projects, err := bootstrap.ProjectService.GetByWorkspaceId(context.Background(), workspace.GetId())
	if err != nil {
		panic(err)
	}

	err = bootstrap.RunnerService.Run(context.Background(), workspace, projects)
	if err != nil {
		panic(err)
	}
}
