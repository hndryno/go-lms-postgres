package loan

import (
	"context"

	"github.com/google/uuid"

	"github.com/hndryno/go-lms-postgresql/internal/adapters/postgresql/db"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	params db.CreateLoanParams,
) error {
	return r.queries.CreateLoan(ctx, params)
}

func (r *Repository) Get(
	ctx context.Context,
	id uuid.UUID,
) (db.Loan, error) {
	return r.queries.GetLoan(ctx, id)
}
