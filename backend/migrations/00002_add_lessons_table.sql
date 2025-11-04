-- +goose Up
-- +goose StatementBegin
CREATE TABLE lessons (
  id SERIAL PRIMARY KEY,
  title VARCHAR(150) NOT NULL,
  language VARCHAR(50) NOT NULL,
  difficulty VARCHAR(50) DEFAULT 'beginner',
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lessons;
-- +goose StatementEnd
