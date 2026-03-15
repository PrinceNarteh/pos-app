-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_returns (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  code VARCHAR(255) NOT NULL,
  NOTE TEXT,
  order_id INT NOT NULL,
  user_id INT NOT NULL,
  date TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

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
