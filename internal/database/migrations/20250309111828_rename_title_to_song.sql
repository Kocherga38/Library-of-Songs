-- +goose Up
ALTER TABLE songs RENAME COLUMN title to "song";

-- +goose Down
ALTER TABLE songs RENAME COLUMN "song" to title;
