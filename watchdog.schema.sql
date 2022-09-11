CREATE TABLE IF NOT EXISTS userver
(
    id             TEXT PRIMARY KEY,
    name           TEXT UNIQUE,
    ip             TEXT,
    port           TEXT,
    protocol       TEXT,
    check_interval INTEGER,
    timeout        INTEGER
);
