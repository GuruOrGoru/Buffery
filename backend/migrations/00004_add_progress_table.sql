-- +goose Up
-- +goose StatementBegin
CREATE TABLE progress (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  lesson_id INT REFERENCES lessons(id),
  status VARCHAR(50) DEFAULT 'incomplete',
  score INT DEFAULT 0,
  updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE progress;
-- +goose StatementEnd
