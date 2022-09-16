CREATE TABLE IF NOT EXISTS userver ( id TEXT PRIMARY KEY, hostname TEXT, alias TEXT, ip TEXT, port INTEGER, protocol TEXT, check_interval INTEGER, timeout INTEGER);
