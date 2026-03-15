-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_return_details (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  product_id UUID NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  qty INT NOT NULL,
  total_price DECIMAL(10,2) NOT NULL,
  order_return_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_order_return_details_order_return 
    FOREIGN KEY(order_return_id)
    REFERENCES order_returns(id) ON DELETE CASCADE,
  CONSTRAINT fk_order_return_details_product 
    FOREIGN KEY(product_id) 
    REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_return_details;
-- +goose StatementEnd
