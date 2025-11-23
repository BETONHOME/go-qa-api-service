-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.answer RENAME TO answers;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE answers RENAME TO answer;
-- +goose StatementEnd
