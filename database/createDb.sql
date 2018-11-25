CREATE TABLE users (id integer PRIMARY KEY, username STRING NOT NULL);
CREATE TABLE files (id integer PRIMARY KEY, filename STRING NOT NULL, data BLOB, owner integer, FOREIGN KEY(owner) REFERENCES users(id));
CREATE TABLE revisions (file integer, number integer, timestamp datetime default CURRENT_TIMESTAMP, FOREIGN KEY(file) references files(id), PRIMARY KEY(file, number));
CREATE TABLE changes (file integer, rev_number integer, position number, character CHARACTER, FOREIGN KEY(file) references revisions(file), FOREIGN KEY(rev_number) references revisions(number), PRIMARY KEY(file, rev_number, position));
