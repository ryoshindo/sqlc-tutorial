-- +migrate Up
CREATE TABLE account_passwords (
    id VARCHAR(32) PRIMARY KEY,
    account_id VARCHAR(32) NOT NULL REFERENCES accounts ON UPDATE NO ACTION ON DELETE CASCADE,
    "hash" BYTEA NOT NULL,
	salt BYTEA NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX ON account_passwords(account_id);

-- +migrate Down
DROP TABLE account_passwords;
