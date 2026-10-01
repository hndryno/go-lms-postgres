package customer

import (
	"context"

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
	params db.CreateCustomerParams,
) (db.Customer, error) {
	return r.queries.CreateCustomer(ctx, params)
}