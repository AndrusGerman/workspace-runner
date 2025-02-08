package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AndrusGerman/go-criteria"
	"github.com/AndrusGerman/workspace-runner/cmd/cmd/bootstrap"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
)

func init() {
	flag.Parse()
}

func main() {
	var bootstrap, err = bootstrap.NewBootstrap()
	if err != nil {
		panic(err)
	}

	if len(flag.Args()) == 0 {
		log.Println("No flag provided")
		return
	}

	if flag.Args()[0] == "run" {
		if len(flag.Args()) < 2 {
			log.Println("No workspace name provided")
			return
		}
		var workspaceName = flag.Args()[1]

		run(bootstrap, workspaceName)
		return
	}

	if flag.Args()[0] == "list" {
		list(bootstrap)
		return
	}

}

func list(bootstrap *bootstrap.Bootstrap) {
	workspaces, err := bootstrap.WorkspaceService.Search(context.Background(), criteria.EmptyCriteria())
	if err != nil {
		panic(err)
	}

	for _, workspace := range workspaces {
		fmt.Printf("- %s\n", workspace.Name)
		listProjects(bootstrap, workspace)
	}
}

func listProjects(bootstrap *bootstrap.Bootstrap, workspace *models.Workspace) {
	projects, err := bootstrap.ProjectService.GetByWorkspaceId(context.Background(), workspace.GetId())
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Printf("  - %s\n", project.Name)
	}

}

func run(bootstrap *bootstrap.Bootstrap, workspaceName string) {
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
