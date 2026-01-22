-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_details (
  id SERIAL PRIMARY KEY,
  product_id INT NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  qty INT NOT NULL,
  total_price DECIMAL(10,2) NOT NULL,
  note TEXT,
  order_id INT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
