package ports

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"
)

// ProjectRepository ...
type ProjectRepository interface {
	BaseRepository[models.Project]
}

type ProjectService interface {
	BaseService[models.Project]
	GetByWorkspaceId(ctx context.Context, workspaceId types.Id) ([]*models.Project, error)
}
