-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(255) NOT NULL,
  bar_code VARCHAR(255),
  image TEXT,
  url TEXT,
  qty INT NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  category_id INT NOT NULL,
  supplier_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
