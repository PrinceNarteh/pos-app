-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_returns (
  id SERIAL PRIMARY KEY,
  code VARCHAR(255) NOT NULL,
  NOTE TEXT,
  order_id INT NOT NULL,
  user_id INT NOT NULL,
  date TIMESTAMP NOT NULL,

  CONSTRAINT fk_order_returns_order 
    FOREIGN KEY(order_id) 
    REFERENCES orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_order_returns_user 
    FOREIGN KEY(user_id) 
    REFERENCES users(id) ON DELETE CASCADE
) 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_returns;
-- +goose StatementEnd
