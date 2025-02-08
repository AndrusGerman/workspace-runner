package ports

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
)

// WorkspaceRepository ...
type WorkspaceRepository interface {
	BaseRepository[models.Workspace]
}

type WorkspaceService interface {
	BaseService[models.Workspace]
	GetByName(ctx context.Context, name string) (*models.Workspace, error)
}
