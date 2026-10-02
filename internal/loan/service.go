package loan

import (
	"context"
	"errors"
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/hndryno/go-lms-postgresql/internal/adapters/postgresql/db"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func numericFromFloat64(value float64) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(int64(value * 100)),
		Exp:   -2,
		Valid: true,
	}
}

func (s *Service) Create(
	ctx context.Context,
	req CreateLoanRequest,
) (db.Loan, error) {
	if req.CustomerID == uuid.Nil {
		return db.Loan{}, ErrCustomerRequired
	}

	if req.PrincipalAmount <= 0 {
		return db.Loan{}, ErrPrincipalInvalid
	}

	if req.InterestRate < 0 {
		return db.Loan{}, ErrInterestInvalid
	}

	if req.Tenor <= 0 {
		return db.Loan{}, ErrTenorInvalid
	}

	loanID := uuid.New()

	err := s.repository.Create(
		ctx,
		db.CreateLoanParams{
			LoanID:          loanID,
			CustomerID:      req.CustomerID,
			PrincipalAmount: numericFromFloat64(req.PrincipalAmount),
			InterestRate:    numericFromFloat64(req.InterestRate),
			Tenor:           req.Tenor,
		},
	)
	if err != nil {
		return db.Loan{}, err
	}

	loan, err := s.repository.Get(ctx, loanID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.Loan{}, errors.New(
				"loan created but failed to retrieve",
			)
		}

		return db.Loan{}, err
	}

	return loan, nil
}

func (s *Service) Get(
	ctx context.Context,
	id uuid.UUID,
) (db.Loan, error) {
	loan, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.Loan{}, ErrLoanNotFound
		}

		return db.Loan{}, err
	}

	return loan, nil
}

func (s *Service) List(
	ctx context.Context,
	req ListLoanRequest,
) ([]db.Loan, int64, error) {
	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	loans, err := s.repository.List(
		ctx,
		int32(req.Limit),
		int32(req.Offset),
	)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repository.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return loans, total, nil
}
