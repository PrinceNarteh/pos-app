-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_return_details (
  id SERIAL PRIMARY KEY,
  product_id INT NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  qty INT NOT NULL,
  total_price DECIMAL(10,2) NOT NULL,
  order_return_id INT NOT NULL,

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
