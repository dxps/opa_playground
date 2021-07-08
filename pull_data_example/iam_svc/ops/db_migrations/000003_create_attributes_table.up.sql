CREATE TABLE IF NOT EXISTS attributes (
    iid            bigserial          PRIMARY KEY,
    owner_id       bigserial          NOT NULL,
    owner_type     smallint           NOT NULL, -- 1 = subject
    name           text               NOT NULL,
    value          text               NOT NULL
);
