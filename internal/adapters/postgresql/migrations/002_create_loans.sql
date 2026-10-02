-- +goose Up

CREATE TABLE loans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    customer_id UUID NOT NULL,
    principal_amount NUMERIC(15, 2) NOT NULL,
    interest_rate NUMERIC(5, 2) NOT NULL,
    tenor INT NOT NULL,

    interest_amount NUMERIC(15, 2) NOT NULL,
    total_amount NUMERIC(15, 2) NOT NULL,
    paid_amount NUMERIC(15, 2) NOT NULL DEFAULT 0,

    status VARCHAR(20) NOT NULL DEFAULT 'active',
    start_date DATE NOT NULL DEFAULT CURRENT_DATE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_loans_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE RESTRICT,

    CONSTRAINT loans_principal_amount_positive
        CHECK (principal_amount > 0),

    CONSTRAINT loans_interest_rate_valid
        CHECK (interest_rate >= 0),

    CONSTRAINT loans_tenor_positive
        CHECK (tenor > 0),

    CONSTRAINT loans_status_valid
        CHECK (status IN ('active', 'paid'))
);

-- +goose Down

DROP TABLE IF EXISTS loans;