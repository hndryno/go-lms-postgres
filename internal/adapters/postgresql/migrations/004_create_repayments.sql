-- +goose Up

CREATE TABLE repayments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    loan_id UUID NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    payment_date DATE NOT NULL DEFAULT CURRENT_DATE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_repayments_loan
        FOREIGN KEY (loan_id)
        REFERENCES loans(id)
        ON DELETE RESTRICT,

    CONSTRAINT repayments_amount_positive
        CHECK (amount > 0)
);

-- +goose Down

DROP TABLE IF EXISTS repayments;