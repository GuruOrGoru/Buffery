-- +goose Up
-- +goose StatementBegin
CREATE TABLE quizzes (
  id SERIAL PRIMARY KEY,
  lesson_id INT REFERENCES lessons(id) ON DELETE CASCADE,
  question TEXT NOT NULL,
  options TEXT[] NOT NULL,
  correct_answer TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE quizzes;
-- +goose StatementEnd
