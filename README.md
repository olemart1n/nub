plan

1. Set up Postgres

2. Make users table & login/signup pages. Just plain HTML forms + bcrypt hashing.

3. Integrate secure Bunny.net direct upload.

4. Save uploaded posts in DB (title, location, tags).

5. Build profile pages, comments, and search last.

### Guiding principles for naming handler files..

1. Group by feature or domain, not by HTTP method.

2. Use suffixes or prefixes to clarify purpose (e.g. \_view, \_api, \_partial).

3. Keep names short but descriptive—they should reflect what the handler does, not how it’s implemented.

### Suggested naming patterns

_pages/views_

- view_index.go
- view_signup.go
- view_dashboard.go

_HTMX parials_

- partial_user.go
- partial_notification.go
- partial_search.go

_API responses (json)_

- api_search.go
- api_upload.go
- api_auth.go

_Auth and session_

- auth.go -> login/logout
- session.go -> session management

---

Needed for application to run

```
brew install redis
brew install postgresql@17
brew services start redis
brew services start postgresql@17
```

---

-- USERS table stays the same
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username TEXT UNIQUE NOT NULL,
password_hash TEXT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- POSTS table: remove image_url
CREATE TABLE posts (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id) ON DELETE CASCADE,
title TEXT NOT NULL,
location TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- IMAGES table: new
CREATE TABLE images (
id SERIAL PRIMARY KEY,
post_id INT REFERENCES posts(id) ON DELETE CASCADE,
image_url TEXT NOT NULL
);

-- TAGS table stays the same
CREATE TABLE tags (
id SERIAL PRIMARY KEY,
name TEXT UNIQUE NOT NULL
);

-- POST_TAGS table stays the same
CREATE TABLE post_tags (
post_id INT REFERENCES posts(id) ON DELETE CASCADE,
tag_id INT REFERENCES tags(id) ON DELETE CASCADE
);

-- COMMENTS table stays the same
CREATE TABLE comments (
id SERIAL PRIMARY KEY,
post_id INT REFERENCES posts(id) ON DELETE CASCADE,
user_id INT REFERENCES users(id) ON DELETE SET NULL,
content TEXT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ENVIRONMENT VARIABLES

```
BUNNY_STORAGE_ZONE=nubglobal
BUNNY_STORAGE_HOST=
BUNNY_UPLOAD_KEY=
BUNNY_PULL_ZONE=

DATABASE_URL=
```
