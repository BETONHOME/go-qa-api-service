-- +goose Up
-- +goose StatementBegin
INSERT INTO answers (question_id, user_id, text, created_at) VALUES
(1, 'boomba', 'test1', NOW()),
(2, 'boomba', 'test2', NOW()),
(2, 'boomba', 'test2.2', NOW()),
(3, 'boomba', 'test3', NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM answers;
-- +goose StatementEnd
