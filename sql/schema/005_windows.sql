-- +goose Up
CREATE TABLE windows (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  wm_class TEXT UNIQUE NOT NULL,
  is_active INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(wm_class) REFERENCES wmclass(wm_class)
);

-- +goose Down
DROP TABLE windows;
