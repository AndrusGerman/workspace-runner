package repository

import (
	mongodb "github.com/AndrusGerman/workspace-runner/internal/adapters/storage/mongo"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"
)

func NewProjectRepository(mongoService *mongodb.Mongo) ports.ProjectRepository {
	return &ProjectRepository{
		BaseRepository: newBaseRepository[models.Project](mongoService, "projects"),
	}
}

type ProjectRepository struct {
	ports.BaseRepository[models.Project]
}
