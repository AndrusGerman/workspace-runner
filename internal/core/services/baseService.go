package services

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"
	"github.com/AndrusGerman/workspace-runner/internal/core/ports"

	"github.com/AndrusGerman/go-criteria"
)

func newBaseService[T types.IBase](repository ports.BaseRepository[T]) ports.BaseService[T] {
	return &BaseService[T]{
		repository: repository,
	}
}

type BaseService[T types.IBase] struct {
	repository ports.BaseRepository[T]
}

func (s *BaseService[T]) GetById(ctx context.Context, id types.Id) (*T, error) {
	return s.repository.GetById(ctx, id)
}

func (s *BaseService[T]) Search(ctx context.Context, filter criteria.Criteria) ([]*T, error) {
	return s.repository.Search(ctx, filter)
}

func (s *BaseService[T]) Delete(ctx context.Context, id types.Id) error {
	return s.repository.Delete(ctx, id)
}

func (s *BaseService[T]) Create(ctx context.Context, element *T) error {
	return s.repository.Create(ctx, element)
}

func (s *BaseService[T]) Update(ctx context.Context, element *T) error {
	return s.repository.Update(ctx, element)
}
