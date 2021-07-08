CREATE TABLE IF NOT EXISTS subjects (
    iid            bigserial                   PRIMARY KEY,
    eid            uuid                        UNIQUE NOT NULL,
    created_at     timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name           text                        NOT NULL,
    email          citext                      UNIQUE NOT NULL,
    password_hash  bytea                       NOT NULL,
    active         bool                        NOT NULL,
    version        integer                     NOT NULL DEFAULT 1
);
