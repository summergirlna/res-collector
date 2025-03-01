CREATE TABLE cpu_stats
(
    serial_number TEXT PRIMARY KEY,
    year          INTEGER NOT NULL,
    month         INTEGER NOT NULL,
    day           INTEGER NOT NULL,
    hour          INTEGER NOT NULL,
    minute        INTEGER NOT NULL,
    second        INTEGER NOT NULL,
    user_time     REAL    NOT NULL,
    system_time   REAL    NOT NULL,
    iowait_time   REAL    NOT NULL
);