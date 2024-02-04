-- +goose Up
-- +goose StatementBegin
INSERT INTO events (title, start_time, duration, description, user_id, remind_for, notified)
VALUES ('Event 1', '2024-01-01 00:01:01', '1 hour', 'Description 1', '6beec901-28a3-h315-b0e0-4d393b83323a', '1 hour', FALSE),
       ('Event 2', '2024-01-02 00:01:01', '2 hours', 'Description 2', '7beec901-48a3-4315-a0e0-5d393b83323b', '2 hours', FALSE),
       ('Event 3', '2024-01-03 00:01:01', '3 hours', 'Description 3', '49dc6fa8-890c-4803-99b7-728a492ea999', '3 hours', FALSE),
       ('Event 4', '2024-01-04 00:01:01', '4 hours', 'Description 4', 'd4fa46c5-0423-4e6c-9f74-9a9405e58fb6', '4 hours', FALSE),
       ('Event 5', '2024-01-05 00:01:01', '5 hours', 'Description 5', '6f16fac2-be54-4778-a2b0-07ac8435afa9', '5 hours', FALSE),
       ('Event 6', '2024-01-06 00:01:01', '6 hours', 'Description 6', 'c5393f09-552e-43d6-b6a4-b377bc2df05e', '6 hours', FALSE),
       ('Event 7', '2024-01-07 00:01:01', '7 hours', 'Description 7', '81603e8b-f525-4c7a-9dd4-1458f73099b9', '7 hours', FALSE),
       ('Event 8', '2024-01-08 00:01:01', '8 hours', 'Description 8', 'c640904f-50f5-4631-8af9-d3c5a2e90733', '8 hours', FALSE);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
