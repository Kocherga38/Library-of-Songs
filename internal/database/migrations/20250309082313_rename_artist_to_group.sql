-- +goose Up
ALTER TABLE songs RENAME COLUMN artist to "group";

-- +goose Down
ALTER TABLE songs RENAME COLUMN "group" to artist;
