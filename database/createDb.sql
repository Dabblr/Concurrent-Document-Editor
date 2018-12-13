CREATE TABLE users (
    id integer PRIMARY KEY,     -- alias for SQLite's ROW_ID column
    username STRING NOT NULL
);

CREATE TABLE files (
    id integer PRIMARY KEY,     -- alias for SQLite's ROW_ID column
    filename STRING NOT NULL,
    data BLOB,
    owner integer,
    FOREIGN KEY(owner) REFERENCES users(id)
);

CREATE TABLE revisions (
    file integer,
    rev_number integer,
    timestamp datetime default CURRENT_TIMESTAMP,
    FOREIGN KEY(file) references files(id),
    PRIMARY KEY(file, rev_number)
);

CREATE TABLE changes (
    id integer PRIMARY KEY,     -- alias for SQLite's ROW_ID column
    file integer,
    rev_number integer,
    position number,
    character CHARACTER,
    FOREIGN KEY(file) references revisions(file),
    FOREIGN KEY(rev_number) references revisions(rev_number)
);