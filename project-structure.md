project-root/
|
|── assets/ # Static assets (images, CSS, JS)
| ├── css/
│ ├── js/
│ └── img/ # Optional: local placeholder images
│
├── templates/ # HTML templates for rendering
│ ├── base.html # Common layout
│ ├── index.html
│ ├── detail.html
│ └── partials/ # Fragments for htmx swaps
│
├── internal/ # Core application logic
│ ├── db/ # DB layer: queries, migrations
│ │ └── postgres.go
│ ├── handlers/ # HTTP handlers
│ │ ├── home.go
│ │ └── media.go
│ ├── bunny/ # Wrapper for Uploadcare API
│ │ ├── client.go
│ │ └── utils.go
│ └── middleware/ # Custom middlewares (auth, logging)
│
├── static/ # Optional: serve static files directly
│
├── public/ # Render static site (if using Static Site option)
│ └── index.html # Required entry point
│
├── go.mod # Go module file
├── go.sum
└── cmd/
|--- main.go # Entry point: setup router, server

1. Set up a postgresql database

2. Find out how to organise file uploads.
