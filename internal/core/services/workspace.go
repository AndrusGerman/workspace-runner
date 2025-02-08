package services

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain"
	"github.com/AndrusGerman/workspace-runner/internal/core/domain/models"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"

	"github.com/AndrusGerman/go-criteria"
)

func NewWorkspaceService(repository ports.WorkspaceRepository) ports.WorkspaceService {
	return &WorkspaceService{
		BaseService: newBaseService[models.Workspace](repository),
	}
}

type WorkspaceService struct {
	ports.BaseService[models.Workspace]
}

func (s *WorkspaceService) GetByName(ctx context.Context, name string) (*models.Workspace, error) {
	var criteriaFilter = criteria.NewFilter("name", criteria.EQUAL, criteria.NewFilterValue(name))
	var criteria = criteria.NewCriteriaBuilder().Filters(criteria.NewFilters([]criteria.Filter{criteriaFilter})).MustGetCriteria()

	var projects, err = s.Search(ctx, criteria)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, domain.ErrNotFound
	}

	return projects[0], nil

}
