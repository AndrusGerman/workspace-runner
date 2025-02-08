package ports

import (
	"context"

	"github.com/AndrusGerman/workspace-runner/internal/core/domain/types"

	criteria "github.com/AndrusGerman/go-criteria"
)

// BaseRepository ...
type BaseRepository[T types.IBase] interface {
	GetById(ctx context.Context, id types.Id) (*T, error)
	Search(ctx context.Context, filter criteria.Criteria) ([]*T, error)
	Delete(ctx context.Context, id types.Id) error
	Create(ctx context.Context, element *T) error
	Update(ctx context.Context, element *T) error
}

type BaseService[T types.IBase] interface {
	GetById(ctx context.Context, id types.Id) (*T, error)
	Search(ctx context.Context, filter criteria.Criteria) ([]*T, error)
	Delete(ctx context.Context, id types.Id) error
	Create(ctx context.Context, element *T) error
	Update(ctx context.Context, element *T) error
}
