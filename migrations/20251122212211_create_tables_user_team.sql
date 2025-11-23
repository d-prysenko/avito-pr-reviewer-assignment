-- +goose Up
CREATE TABLE IF NOT EXISTS revass.team
(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(128) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS revass.users
(
    id varchar(32) PRIMARY KEY,
    username varchar(128) NOT NULL,
    is_active boolean NOT NULL,
    team_id integer NOT NULL REFERENCES revass.team(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE revass.users;
DROP TABLE revass.team;
