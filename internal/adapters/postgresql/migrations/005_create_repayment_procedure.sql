-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE create_repayment(
    IN p_repayment_id UUID,
    IN p_loan_id UUID,
    IN p_amount NUMERIC(15, 2)
)
LANGUAGE plpgsql
AS $$
DECLARE
    v_total_amount NUMERIC(15, 2);
    v_paid_amount NUMERIC(15, 2);
    v_remaining_amount NUMERIC(15, 2);
BEGIN

    /*
     * Lock loan row.
     * Ini penting supaya dua repayment
     * tidak mengubah paid_amount secara bersamaan.
     */
    SELECT
        total_amount,
        paid_amount
    INTO
        v_total_amount,
        v_paid_amount
    FROM loans
    WHERE id = p_loan_id
    FOR UPDATE;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'loan not found';
    END IF;

    IF p_amount <= 0 THEN
        RAISE EXCEPTION 'repayment amount must be greater than zero';
    END IF;

    v_remaining_amount :=
        v_total_amount - v_paid_amount;

    IF p_amount > v_remaining_amount THEN
        RAISE EXCEPTION 'repayment amount exceeds remaining loan amount';
    END IF;

    INSERT INTO repayments (
        id,
        loan_id,
        amount,
        payment_date
    )
    VALUES (
        p_repayment_id,
        p_loan_id,
        p_amount,
        CURRENT_DATE
    );

    UPDATE loans
    SET
        paid_amount = paid_amount + p_amount,
        status = CASE
            WHEN paid_amount + p_amount >= total_amount
                THEN 'paid'
            ELSE 'active'
        END,
        updated_at = NOW()
    WHERE id = p_loan_id;

END;
$$;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP PROCEDURE IF EXISTS create_repayment(
    UUID,
    UUID,
    NUMERIC
);
-- +goose StatementEnd