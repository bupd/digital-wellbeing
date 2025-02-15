-- +goose Up
CREATE TABLE keys (
  id TEXT PRIMARY KEY,
  keyname TEXT NOT NULL,
  keycode INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE keys;
