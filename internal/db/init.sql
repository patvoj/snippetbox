BEGIN;

-- 1. Create 'web' user if it doesn't exist.
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'web') THEN
        CREATE ROLE web LOGIN PASSWORD 'securepassword123';
    END IF;
END
$$;

-- 2. Create the snippets table. The 'SERIAL PRIMARY KEY' will implicitly create 'snippets_id_seq'.
CREATE TABLE IF NOT EXISTS snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires TIMESTAMPTZ NOT NULL
);

-- 3. Create an index for the 'created' column.
CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);

-- 4. Grant privileges on *existing* table 'snippets' to 'web' user.
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE snippets TO web;

-- 5. Grant privileges on the *existing* sequence 'snippets_id_seq' to 'web' user.
GRANT USAGE, SELECT ON SEQUENCE snippets_id_seq TO web;

-- 6. Set default privileges for 'web' role on *future* tables created in 'public' schema.
ALTER DEFAULT PRIVILEGES FOR ROLE web IN SCHEMA public
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO web;

-- 7. Set default privileges for 'web' role on *future* sequences created in 'public' schema.
ALTER DEFAULT PRIVILEGES FOR ROLE web IN SCHEMA public
GRANT USAGE, SELECT ON SEQUENCES TO web;

-- 8. Seed data into the snippets table.
INSERT INTO snippets (title, content, created, expires) VALUES
('An old silent pond', 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', NOW(), NOW() + INTERVAL '365 days'),
('Over the wintry forest', 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', NOW(), NOW() + INTERVAL '365 days'),
('First autumn morning', 'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', NOW(), NOW() + INTERVAL '7 days');

COMMIT;
