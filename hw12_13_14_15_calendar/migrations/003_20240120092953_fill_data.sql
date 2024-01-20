-- +goose Up
-- +goose StatementBegin
INSERT INTO events (title, start_time, duration, description, user_id, remind_for, notified)
VALUES ('Event 1', '2024-01-01 00:01:01', '1 hour', 'Description 1', gen_random_uuid(), '1 hour', FALSE),
       ('Event 2', '2024-01-02 00:01:01', '2 hours', 'Description 2', gen_random_uuid(), '2 hours', FALSE),
       ('Event 3', '2024-01-03 00:01:01', '3 hours', 'Description 3', gen_random_uuid(), '3 hours', FALSE),
       ('Event 4', '2024-01-04 00:01:01', '4 hours', 'Description 4', gen_random_uuid(), '4 hours', FALSE),
       ('Event 5', '2024-01-05 00:01:01', '5 hours', 'Description 5', gen_random_uuid(), '5 hours', FALSE),
       ('Event 6', '2024-01-06 00:01:01', '6 hours', 'Description 6', gen_random_uuid(), '6 hours', FALSE),
       ('Event 7', '2024-01-07 00:01:01', '7 hours', 'Description 7', gen_random_uuid(), '7 hours', FALSE),
       ('Event 8', '2024-01-08 00:01:01', '8 hours', 'Description 8', gen_random_uuid(), '8 hours', FALSE);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
