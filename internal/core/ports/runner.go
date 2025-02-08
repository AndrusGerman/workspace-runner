package ports

import (
	"context"
	"io"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
)

type RunnerService interface {
	Run(ctx context.Context, workspace *models.Workspace, projects []*models.Project) error
}

type RunnerLogger interface {
	GetStdout(projectName string) io.Writer
	GetStderr(projectName string) io.Writer
}
