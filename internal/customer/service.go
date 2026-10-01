package customer

import (
	"context"
	"errors"
	"strings"

	"github.com/hndryno/go-lms-postgresql/internal/adapters/postgresql/db"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrNameRequired  = errors.New("name is required")
	ErrEmailRequired = errors.New("email is required")
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
