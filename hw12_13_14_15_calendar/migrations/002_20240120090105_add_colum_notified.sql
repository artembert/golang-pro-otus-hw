-- +goose Up
-- +goose StatementBegin
ALTER TABLE events
    ADD COLUMN notified BOOL NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE events
    DROP COLUMN notified;
-- +goose StatementEnd
