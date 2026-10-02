package repayment

import (
	"context"
	"math/big"

	"github.com/google/uuid"
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

func (s *Service) Create(
	ctx context.Context,
	req CreateRepaymentRequest,
) error {
	if req.LoanID == uuid.Nil {
		return ErrLoanRequired
	}

	if req.Amount <= 0 {
		return ErrAmountInvalid
	}

	repaymentID := uuid.New()

	return s.repository.Create(
		ctx,
		db.CreateRepaymentParams{
			RepaymentID: repaymentID,
			LoanID:      req.LoanID,
			Amount: pgtype.Numeric{
				Int:   big.NewInt(int64(req.Amount)),
				Exp:   0,
				Valid: true,
			},
		},
	)
}
