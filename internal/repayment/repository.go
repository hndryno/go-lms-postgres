package repayment

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
	params db.CreateRepaymentParams,
) error {
	return r.queries.CreateRepayment(ctx, params)
}

func (r *Repository) Get(
	ctx context.Context,
	id uuid.UUID,
) (db.Repayment, error) {
	return r.queries.GetRepayment(ctx, id)
}

func (r *Repository) List(
	ctx context.Context,
	loanID uuid.UUID,
	limit int32,
	offset int32,
) ([]db.Repayment, error) {
	return r.queries.ListRepayments(ctx, db.ListRepaymentsParams{
		LoanID: loanID,
		Limit:  limit,
		Offset: offset,
	})
}

func (r *Repository) Count(
	ctx context.Context,
	loanID uuid.UUID,
) (int64, error) {
	return r.queries.CountRepayments(ctx, loanID)
}
