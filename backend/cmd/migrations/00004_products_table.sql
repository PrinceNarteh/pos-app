-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  code VARCHAR(255) NOT NULL,
  bar_code VARCHAR(255),
  image TEXT,
  url TEXT,
  qty INT NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  category_id UUID NOT NULL,
  supplier_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_products_category
    FOREIGN KEY (category_id)
    REFERENCES categories(id) 
    ON DELETE CASCADE,
  CONSTRAINT fk_products_supplier
    FOREIGN KEY (supplier_id)
    REFERENCES suppliers(id) 
    ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
