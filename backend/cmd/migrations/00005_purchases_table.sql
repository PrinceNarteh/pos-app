-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS purchases (
  id SERIAL PRIMARY KEY,
  code VARCHAR(255) NOT NULL,
  note TEXT,
  total DECIMAL(10,2) NOT NULL,
  ppn INT NOT NULL,
  grand_total DECIMAL(10, 2) NOT NULL,
  user_id INT NOT NULL,
  date TIMESTAMP NOT NULL,

  CONSTRAINT fk_purchases_user 
    FOREIGN KEY(user_id) 
    REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS purchases;
-- +goose StatementEnd
