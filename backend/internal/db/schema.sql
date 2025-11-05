
-- Users
CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    username      TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    email         TEXT UNIQUE
);

-- Posts
CREATE TABLE posts (
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title       TEXT NOT NULL,
    location    TEXT,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

--Images
CREATE TABLE images (
    id          SERIAL PRIMARY KEY,
    post_id     INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    image_url   TEXT NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- Tags

CREATE TABLE tags (
    id    SERIAL PRIMARY KEY,
    name  TEXT NOT NULL UNIQUE
);

-- Post <-> Tag relationship
CREATE TABLE post_tags (
    post_id INTEGER NOT NULL,
    tag_id  INTEGER NOT NULL,
    PRIMARY KEY (post_id, tag_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Comments
CREATE TABLE comments (
    id          SERIAL PRIMARY KEY,
    post_id     INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    user_id     INTEGER REFERENCES users(id) ON DELETE SET NULL,
    content     TEXT NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
)


-- Index for tag search
CREATE INDEX idx_tags_name ON tags(name);

-- Index for location search
CREATE INDEX idx_posts_location ON posts(location);

-- Full-text search index
CREATE INDEX idx_posts_search 
ON posts USING GIN(to_tsvector('english', title || ' ' || location));
