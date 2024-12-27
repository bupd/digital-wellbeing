-- +goose Up
CREATE TABLE mouse (
  id TEXT PRIMARY KEY,
  keyname TEXT UNIQUE NOT NULL,
  is_moved INTEGER NOT NULL,
  is_scroll INTEGER NOT NULL,
  distance INTEGER NOT NULL,
  keycode INTEGER NOT NULL,
  x INTEGER,
  y INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE mouse;