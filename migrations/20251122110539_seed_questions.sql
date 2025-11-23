-- +goose Up
-- +goose StatementBegin
INSERT INTO questions (text, created_at) VALUES
('test quest 1', NOW()),
('test quest 2', NOW()),
('test quest 3', NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM questions;
-- +goose StatementEnd
