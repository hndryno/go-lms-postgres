-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE create_loan(
    IN p_loan_id UUID,
    IN p_customer_id UUID,
    IN p_principal_amount NUMERIC(15, 2),
    IN p_interest_rate NUMERIC(5, 2),
    IN p_tenor INT
)
LANGUAGE plpgsql
AS $$
DECLARE
    v_interest_amount NUMERIC(15, 2);
    v_total_amount NUMERIC(15, 2);
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM customers
        WHERE id = p_customer_id
    ) THEN
        RAISE EXCEPTION 'customer not found';
    END IF;

    IF p_principal_amount <= 0 THEN
        RAISE EXCEPTION 'principal amount must be greater than zero';
    END IF;

    IF p_interest_rate < 0 THEN
        RAISE EXCEPTION 'interest rate cannot be negative';
    END IF;

    IF p_tenor <= 0 THEN
        RAISE EXCEPTION 'tenor must be greater than zero';
    END IF;

    v_interest_amount :=
        p_principal_amount * p_interest_rate / 100;

    v_total_amount :=
        p_principal_amount + v_interest_amount;

    INSERT INTO loans (
        id,
        customer_id,
        principal_amount,
        interest_rate,
        tenor,
        interest_amount,
        total_amount,
        paid_amount,
        status,
        start_date
    )
    VALUES (
        p_loan_id,
        p_customer_id,
        p_principal_amount,
        p_interest_rate,
        p_tenor,
        v_interest_amount,
        v_total_amount,
        0,
        'active',
        CURRENT_DATE
    );
END;
$$;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP PROCEDURE IF EXISTS create_loan(
    UUID,
    UUID,
    NUMERIC,
    NUMERIC,
    INT
);
-- +goose StatementEnd