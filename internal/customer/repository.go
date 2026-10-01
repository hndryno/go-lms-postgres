package customer

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
	params db.CreateCustomerParams,
) (db.Customer, error) {
	return r.queries.CreateCustomer(ctx, params)
}

func (r *Repository) List(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]db.Customer, error) {
	return r.queries.ListCustomers(
		ctx,
		db.ListCustomersParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}

func (r *Repository) Count(
	ctx context.Context,
) (int64, error) {
	return r.queries.CountCustomers(ctx)
}

func (r *Repository) Get(
	ctx context.Context,
	id uuid.UUID,
) (db.Customer, error) {
	return r.queries.GetCustomer(ctx, id)
}

func (r *Repository) Update(
	ctx context.Context,
	params db.UpdateCustomerParams,
) (db.Customer, error) {
	return r.queries.UpdateCustomer(ctx, params)
}
