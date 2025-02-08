package main

import (
	"context"
	"fmt"

	"github.com/AndrusGerman/workspace-runner/internal/adapters/config"
	mongodb "github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo"
	"github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo/repository"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/services"
)

func main() {

	var err error
	var config = config.NewConfig()

	var mongo *mongodb.Mongo
	mongo, err = mongodb.NewMongo(config)
	if err != nil {
		panic(err)
	}

	var workspaceRepository = repository.NewWorkspaceRepository(mongo)
	var workspaceService = services.NewWorkspaceService(workspaceRepository)

	var projectRepository = repository.NewProjectRepository(mongo)
	var projectService = services.NewProjectService(projectRepository)

	var workspace = models.NewWorkspace("workspace 1", "workspace 1 description")
	err = workspaceService.Create(context.Background(), workspace)
	if err != nil {
		panic(err)
	}

	var project = models.NewProject("project 1", workspace.GetId(), ".", models.NewCmd("ls", []string{"-l"}, []models.Env{}))
	err = projectService.Create(context.Background(), project)
	if err != nil {
		panic(err)
	}

	fmt.Println("project created")

}
