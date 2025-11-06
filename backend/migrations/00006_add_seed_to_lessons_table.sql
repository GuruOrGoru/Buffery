-- +goose Up
-- +goose StatementBegin
INSERT INTO lessons (title, language, difficulty, content, slug, summary)
VALUES
('Introduction to Programming', 'general', 'beginner',
 '# What is Programming?\nProgramming means giving instructions to a computer...',
 'intro-to-programming',
 'Learn what programming is and how computers follow instructions.'),
('Variables in Go', 'go', 'beginner',
 '## Variables in Go\n```go\nvar name string = "Guru"\n```',
 'variables-in-go',
 'Learn variables, data types, and assignment in Go.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM lessons
WHERE slug IN (
  'intro-to-programming',
  'variables-in-go'
);
-- +goose StatementEnd
