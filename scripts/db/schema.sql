CREATE TABLE IF NOT EXISTS cpu_stats
(
    stats_id    TEXT PRIMARY KEY,
    user_time   REAL NOT NULL,
    system_time REAL NOT NULL,
    iowait_time REAL NOT NULL,
    FOREIGN KEY (stats_id) REFERENCES stats_management (stats_id)
);

CREATE TABLE IF NOT EXISTS memory_stats
(
    stats_id         TEXT PRIMARY KEY,
    application_used REAL NOT NULL,
    FOREIGN KEY (stats_id) REFERENCES stats_management (stats_id)

);

CREATE TABLE IF NOT EXISTS stats_management
(
    stats_id TEXT PRIMARY KEY,
    year     INTEGER NOT NULL,
    month    INTEGER NOT NULL,
    day      INTEGER NOT NULL,
    hour     INTEGER NOT NULL,
    minute   INTEGER NOT NULL,
    second   INTEGER NOT NULL
);