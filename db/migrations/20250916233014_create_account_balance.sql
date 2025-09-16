-- +goose Up
CREATE TABLE account_balance (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    account_id BIGINT UNSIGNED NOT NULL,
    balance DECIMAL(12,2) NOT NULL DEFAULT 0,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE account_balance;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
