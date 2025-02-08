package main

import (
	"context"
	"fmt"

	"github.com/AndrusGerman/go-criteria"
	"github.com/AndrusGerman/workspace-runner/cmd/workspace-runner/bootstrap"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
)

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
