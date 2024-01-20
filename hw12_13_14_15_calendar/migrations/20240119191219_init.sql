-- +goose Up
-- +goose StatementBegin
CREATE table events
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    title       text                                       NOT NULL,
    start_time  timestamp                                  NOT NULL,
    duration interval NOT NULL,
    description text,
    user_id     UUID                                       NOT NULL,
    remind_for  TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
