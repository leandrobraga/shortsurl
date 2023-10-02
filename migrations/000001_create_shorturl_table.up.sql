CREATE TABLE IF NOT EXISTS shorturl (
    id bigserial PRIMARY KEY,
    url text NOT NULL,
    code varchar(6) NOT NULL,
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);