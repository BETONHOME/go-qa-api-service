-- +goose Up
-- +goose StatementBegin
CREATE TABLE answer (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT answer_question
    FOREIGN KEY(question_id)
    REFERENCES questions(id)
    ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE answers;
-- +goose StatementEnd