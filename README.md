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

**posts**
Column Type Notes

id SERIAL PK
user_id INT FK → users.id Author
title TEXT User title for image
location TEXT Separate input field
image_url TEXT Bunny.net CDN URL
created_at TIMESTAMP

**tags**
Column Type Notes

id SERIAL PK
name TEXT UNIQUE

**users**
Column Type Notes

id SERIAL PK
username TEXT UNIQUE Required for login
password_hash TEXT Store hashed password, never plain text
created_at TIMESTAMP

**post_tags**
Column Type Notes

post_id INT FK → posts.id
tag_id INT FK → tags.id

**comments**
Column Type Notes

id SERIAL PK
post_id INT FK → posts.id
user_id INT FK → users.id Null allowed for guest comments if needed
content TEXT
created_at TIMESTAMP

ENVIRONMENT VARIABLES

```
BUNNY_STORAGE_ZONE=nubglobal
BUNNY_STORAGE_HOST=
BUNNY_UPLOAD_KEY=
BUNNY_PULL_ZONE=

DATABASE_URL=
```
