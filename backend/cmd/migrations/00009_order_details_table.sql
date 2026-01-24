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
  order_id INT NOT NULL,

  CONSTRAINT fk_order_details_order 
    FOREIGN KEY(order_id) 
    REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_order_details_product 
    FOREIGN KEY(product_id) 
    REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_details;
-- +goose StatementEnd
