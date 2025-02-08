package main

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/adapters/config"
	mongodb "github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo"
	"github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo/repository"
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

	var runnerLogger = services.NewRunnerLogger()
	var runnerService = services.NewRunnerService(runnerLogger)

	workspace, err := workspaceService.GetByName(context.Background(), "workspace:kororo")
	if err != nil {
		panic(err)
	}

	projects, err := projectService.GetByWorkspaceId(context.Background(), workspace.GetId())
	if err != nil {
		panic(err)
	}

	runnerService.Run(context.Background(), workspace, projects)
}
