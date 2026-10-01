package customer

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/hndryno/go-lms-postgresql/internal/adapters/postgresql/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

type CreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (s *Service) Create(
	ctx context.Context,
	req CreateRequest,
) (db.Customer, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if req.Name == "" {
		return db.Customer{}, ErrNameRequired
	}

	if req.Email == "" {
		return db.Customer{}, ErrEmailRequired
	}

	return s.repository.Create(
		ctx,
		db.CreateCustomerParams{
			Name:  req.Name,
			Email: req.Email,
			Phone: pgtype.Text{
				String: req.Phone,
				Valid:  req.Phone != "",
			},
			Address: pgtype.Text{
				String: req.Address,
				Valid:  req.Address != "",
			},
		},
	)
}

func (s *Service) List(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]db.Customer, int64, error) {
	customers, err := s.repository.List(
		ctx,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repository.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (s *Service) Get(
	ctx context.Context,
	id uuid.UUID,
) (db.Customer, error) {
	customer, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.Customer{}, ErrCustomerNotFound
		}

		return db.Customer{}, err
	}

	return customer, nil
}

func (s *Service) Update(
	ctx context.Context,
	id uuid.UUID,
	req UpdateCustomerRequest,
) (db.Customer, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Phone = strings.TrimSpace(req.Phone)
	req.Address = strings.TrimSpace(req.Address)

	if req.Name == "" {
		return db.Customer{}, ErrNameRequired
	}

	if req.Email == "" {
		return db.Customer{}, ErrEmailRequired
	}

	customer, err := s.repository.Update(
		ctx,
		db.UpdateCustomerParams{
			ID:    id,
			Name:  req.Name,
			Email: req.Email,
			Phone: pgtype.Text{
				String: req.Phone,
				Valid:  req.Phone != "",
			},
			Address: pgtype.Text{
				String: req.Address,
				Valid:  req.Address != "",
			},
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.Customer{}, ErrCustomerNotFound
		}

		return db.Customer{}, err
	}

	return customer, nil
}
