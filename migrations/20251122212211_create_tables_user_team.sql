-- +goose Up
CREATE TABLE IF NOT EXISTS revass.users
(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    str_id varchar(128) UNIQUE NOT NULL,
    username varchar(128) NOT NULL,
    is_active boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS revass.team
(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(128) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS revass.team_user
(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    team_id integer NOT NULL REFERENCES revass.team(id) ON DELETE CASCADE,
    user_id integer NOT NULL REFERENCES revass.users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE revass.team_user;
DROP TABLE revass.team;
DROP TABLE revass.users;
