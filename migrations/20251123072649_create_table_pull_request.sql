-- +goose Up
CREATE TYPE PR_STATUS AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS revass.pull_request
(
    id varchar(32) PRIMARY KEY,
    name varchar(128) NOT NULL,
    author_id varchar(32) NOT NULL REFERENCES revass.users(id) ON DELETE CASCADE,
    status PR_STATUS DEFAULT 'OPEN',
    merged_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS revass.pr_reviewer
(
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    pr_id varchar(32) NOT NULL REFERENCES revass.pull_request(id) ON DELETE CASCADE,
    reviewer_id varchar(32) NOT NULL REFERENCES revass.users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE revass.pr_reviewer;
DROP TABLE revass.pull_request;
DROP TYPE PR_STATUS;