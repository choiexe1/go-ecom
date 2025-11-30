-- +goose Up
-- +goose StatementBegin
ALTER TABLE products RENAME COLUMN price_in_centers TO price_cents;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products RENAME COLUMN price_cents TO price_in_centers;
-- +goose StatementEnd
