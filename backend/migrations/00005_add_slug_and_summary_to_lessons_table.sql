-- +goose Up
-- +goose StatementBegin
ALTER TABLE lessons
ADD COLUMN slug VARCHAR(150) UNIQUE,
ADD COLUMN summary TEXT;

UPDATE lessons SET slug = replace(lower(title), ' ', '-') WHERE slug IS NULL;

ALTER TABLE lessons
ALTER COLUMN slug SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE lessons
DROP COLUMN slug,
DROP COLUMN summary;
-- +goose StatementEnd
