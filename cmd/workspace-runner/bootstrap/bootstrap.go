package bootstrap

import (
	"github.com/AndrusGerman/workspace-runner/internal/adapters/config"
	mongodb "github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo"
	"github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo/repository"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
	"github.com/AndrusGerman/workspace-runner/internal/core/services"
)

type Bootstrap struct {
	workspaceRepository ports.WorkspaceRepository
	WorkspaceService    ports.WorkspaceService

	projectRepository ports.ProjectRepository
	ProjectService    ports.ProjectService

	runnerLogger  ports.RunnerLogger
	RunnerService ports.RunnerService

	config *config.Config
	mongo  *mongodb.Mongo
}

func NewBootstrap() (*Bootstrap, error) {

	var err error
	var config = config.NewConfig()

	var mongo *mongodb.Mongo
	mongo, err = mongodb.NewMongo(config)
	if err != nil {
		return nil, err
	}

	var workspaceRepository = repository.NewWorkspaceRepository(mongo)
	var workspaceService = services.NewWorkspaceService(workspaceRepository)

	var projectRepository = repository.NewProjectRepository(mongo)
	var projectService = services.NewProjectService(projectRepository)

	var runnerLogger = services.NewRunnerLogger()
	var runnerService = services.NewRunnerService(runnerLogger)

	return &Bootstrap{
		workspaceRepository: workspaceRepository,
		WorkspaceService:    workspaceService,
		projectRepository:   projectRepository,
		ProjectService:      projectService,
		runnerLogger:        runnerLogger,
		RunnerService:       runnerService,
		config:              config,
		mongo:               mongo,
	}, nil
}
