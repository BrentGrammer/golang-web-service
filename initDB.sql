-- CREATE DATABASE readinglist;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'readinglist') THEN
        CREATE ROLE readinglist WITH LOGIN PASSWORD E'${POSTGRES_PASSWORD}';
    END IF;
END $$;


CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    rating real NOT NULL,
    version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE ON books TO readinglist;

GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;

-- Create role with password from environment variable

