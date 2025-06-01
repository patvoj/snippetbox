BEGIN;

-- Create snippetbox schema and web user
CREATE TABLE IF NOT EXISTS snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);

-- Create 'web' user and grant privileges
DO $$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'web') THEN
      CREATE ROLE web LOGIN PASSWORD 'securepassword123';
   END IF;
END
$$;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE snippets TO web;

-- Seed data
INSERT INTO snippets (title, content, created, expires) VALUES
('An old silent pond', 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', NOW(), NOW() + INTERVAL '365 days'),
('Over the wintry forest', 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', NOW(), NOW() + INTERVAL '365 days'),
('First autumn morning', 'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', NOW(), NOW() + INTERVAL '7 days');

COMMIT;
