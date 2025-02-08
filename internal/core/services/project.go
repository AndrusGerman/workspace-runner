package services

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"

	"github.com/AndrusGerman/go-criteria"
)

func NewProjectService(repository ports.ProjectRepository) ports.ProjectService {
	return &ProjectService{
		BaseService: newBaseService[models.Project](repository),
	}
}

type ProjectService struct {
	ports.BaseService[models.Project]
}

func (s *ProjectService) GetByWorkspaceId(ctx context.Context, workspaceId types.Id) ([]*models.Project, error) {
	var criteriaFilter = criteria.NewFilter("workspace_id", criteria.EQUAL, criteria.NewFilterValue(workspaceId.GetPrimitive()))
	var criteria = criteria.NewCriteriaBuilder().Filters(criteria.NewFilters([]criteria.Filter{criteriaFilter})).MustGetCriteria()
	return s.Search(ctx, criteria)
}
