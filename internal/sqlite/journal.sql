CREATE TABLE IF NOT EXISTS records
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        TEXT,
    name        TEXT,
    description TEXT,
    start       INTEGER,
    end         INTEGER
);