---
name: regenerate-swagger
description: Regenerate the Swagger/OpenAPI docs for todo-list-app after changing handler annotations, DTO structs, or the global @title block on main.go. Use when the user asks to update, regenerate, refresh, or rebuild the Swagger/OpenAPI spec, or after modifying @Summary/@Router/@Param comments, adding a new endpoint, or renaming a request/response type in internal/handler/dto.go. Also use if `go build` or the /swagger/doc.json response looks stale compared to the code.
---

# Regenerate Swagger docs

The `docs/` package (`docs.go`, `swagger.json`, `swagger.yaml`) is **generated** from swag annotations. Regenerate it whenever the annotations, DTOs, or the global `@title` / `@securityDefinitions` block on `main.go` change.

Never hand-edit files under `docs/` — they are 100% regenerable and edits are lost on the next `swag init`.

## Steps

1. **Install swag once**, skip if `swag --version` already works. The binary lands in `$(go env GOPATH)/bin/swag`:
   ```
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. **Regenerate** from the project root (`todo-list-app/`):
   ```
   $(go env GOPATH)/bin/swag init -g main.go -o docs --parseDependency --parseInternal
   ```
   Flags:
   - `-g main.go` — the entry file with the `@title` / `@securityDefinitions` block.
   - `-o docs` — output directory. `main.go` blank-imports it: `_ "todo-list-app/docs"`.
   - `--parseInternal` — required so swag can see types under `internal/` (all our handlers + models live there).
   - `--parseDependency` — required so referenced types from imported packages (e.g. `models.User`) resolve.

3. **Verify**:
   ```
   go build ./...
   go test -run TestSwaggerRoutes ./internal/handler/
   ```
   `TestSwaggerRoutes` fetches `/swagger/doc.json`, `/swagger/index.html`, and confirms the `/swagger` → `/swagger/index.html` redirect. If it fails, either the regeneration didn't happen or an annotation is malformed.

## Common issues

- **`warning: failed to evaluate const mProfCycleWrap ...`** — harmless. Comes from swag scanning the Go stdlib for constants and does not affect the generated spec.
- **A type is missing from the generated schema** — the referenced Go type must be **exported** (capitalized) and reachable via `--parseInternal`. If a handler uses an anonymous struct for its request/response, extract it into a named type in `internal/handler/dto.go` and reference it by name in the `@Param body body <TypeName>` / `@Success 200 {object} <TypeName>` annotation.
- **`/swagger/doc.json` returns an old spec after regeneration** — the app was built from a stale `docs/` package. Rebuild (`go build ./...` or `docker compose up --build`).
- **`swag: command not found`** — `$(go env GOPATH)/bin` isn't on `$PATH`. Always invoke via `$(go env GOPATH)/bin/swag` from within this project, or `export PATH="$(go env GOPATH)/bin:$PATH"` for the session.
