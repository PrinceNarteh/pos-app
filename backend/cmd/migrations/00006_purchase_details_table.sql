-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS purchase_details (
  id SERIAL PRIMARY KEY,
  product_id INT NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  qty INT NOT NULL,
  total_price DECIMAL(10,2) NOT NULL,
  purchase_id INT NOT NULL,

  CONSTRAINT fk_purchase_details_purchase 
    FOREIGN KEY(purchase_id) 
    REFERENCES purchases(id) ON DELETE CASCADE,
  CONSTRAINT fk_purchase_details_product 
    FOREIGN KEY(product_id) 
    REFERENCES products(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS purchase_details;
-- +goose StatementEnd
