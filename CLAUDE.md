# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

SMTP admin UI — a server-rendered Go web app for administering an SMTP mailbox database (domains, mailbox users, mailboxes, sessions). Module path: `github.com/dgb9/smtp-admin` (Go 1.26.2).

The repository is in an **early scaffolding stage**. `main.go` is a stub at the repo root, and the only Go file under `internal/` is an empty `SessionData` struct (`internal/data/session.go`). Most "code" today lives as HTML wireframes in `docs/site/` and SQL DDL in `docs/`. Treat the directory layout in section 2 below as the **target architecture** to grow into — do not assume those packages exist yet.

## Commands

Standard Go toolchain, no build scripts present:

- Build: `go build ./...`
- Run: `go run .` (currently prints a placeholder string)
- Test: `go test ./...`
- Run a single test: `go test ./internal/data -run TestName`
- Vet: `go vet ./...`

## Tech Stack

* **Language:** Go
* **Routing:** Standard library `net/http` (no framework chosen yet)
* **Templating:** Standard library `html/template` (server-rendered HTML)
* **CSS/Frontend:** Plain CSS only — see `docs/site/dep/style.css`. Wireframes use no JS framework; introduce HTMX only when a partial-swap interaction is actually needed.
* **Database:** MariaDB via `database/sql` only. **Do not introduce GORM** or any ORM.

## Target Architecture

Grow new code into this layout. Do not create arbitrary directories.

```text
├── cmd/server/          # Main entry point (move main.go here)
├── internal/
│   ├── database/        # DB connection and SQL queries
│   ├── handlers/        # HTTP handlers / controllers
│   └── models/          # Structs and data models
├── ui/
│   ├── static/          # CSS, JS, images
│   └── html/
│       ├── pages/       # Full view pages
│       └── partials/    # Reusable components (HTMX snippets)
```

Note the mismatch with the current `internal/data/` package — when real code lands, prefer the `database`/`handlers`/`models` split above over extending `internal/data`.

## Domain Model

Schema lives in `docs/ddl.sql` with later changes in `docs/db_changes.sql`. Authoritative tables:

- `domain` — mail domain; supports a catch-all login.
- `mailbox_user` — login/password account scoped to a domain, with `admin_ind` flag (added in `db_changes.sql`) gating admin UI access.
- `mailbox` — per-user mailbox with a wide set of IMAP-style flag columns (`flag_subscribed`, `flag_junk`, `flag_trash`, …).
- `user_session` (in `db_changes.sql`, supersedes the earlier `session` table in `ddl.sql`) — session storage keyed by `session_id`, with `expired_ind` and `session_data` MEDIUMTEXT.

All `*_ind` columns are `VARCHAR(1)` storing `'Y'`/`'N'`, not booleans. Collation is `utf8mb4_general_ci` throughout.

## UI Wireframes

`docs/site/*.html` are static wireframes that define the intended pages and form fields — use them as the spec when implementing templates:

- `login.html`, `chpass.html` — auth flows
- `domains.html`, `viewdomain.html`, `editdomain.html`, `cddomains.html` — domain admin
- `cdusers.html`, `edituser.html` — mailbox user admin

`cd*` prefixes denote create/delete dialogs. Shared styles live in `docs/site/dep/style.css`. When porting a wireframe to a real template, keep the existing class names (`.header`, `.menu`, `.form`, `.form-label`, `.form-input`, `.btn`, `.error`) so the CSS continues to apply.

## Working Style (from project owner)

- **Be lazy** — implement only what the current request needs; do not build features ahead of demand.
- **Be concise** — short answers, no long code snippets or explanations unless asked.
