package repository

import (
	mongodb "github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
)

func NewWorkspaceRepository(mongoService *mongodb.Mongo) ports.WorkspaceRepository {
	return &WorkspaceRepository{
		BaseRepository: newBaseRepository[models.Workspace](mongoService, "workspaces"),
	}
}

type WorkspaceRepository struct {
	ports.BaseRepository[models.Workspace]
}
