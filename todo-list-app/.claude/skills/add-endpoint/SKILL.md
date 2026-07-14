---
name: add-endpoint
description: Add a new HTTP endpoint to todo-list-app end-to-end across all layers (DTO → service → repository → router → Swagger → tests). Use when the user asks to add, expose, or implement a new REST endpoint, extend the API with a new operation on an existing resource, or add a new admin/role-gated action. Follows the layered pattern already established in this codebase — do not invent alternatives (e.g., putting ownership checks in middleware, or returning raw pgx errors from services).
---

# Add a new endpoint

Follow this checklist in order. Each step points to the file(s) to edit and the local convention to preserve.

## 1. DTOs — `internal/handler/dto.go`

- Add a named request struct if the endpoint takes a JSON body:
  ```go
  type XyzRequest struct {
      Field string `json:"field" example:"..."`
  }
  ```
- Add a response struct only if the return shape isn't already `models.*`, `StatusResponse`, or `ErrorResponse`.
- Named types are required — swag can't reference anonymous structs.

## 2. Service — `internal/service/<domain>_service.go`

Signature convention: `func (s *XyzService) DoThing(ctx context.Context, ..., caller Caller) (returnType, error)` for anything that reads or writes user-owned data. Read-only global operations skip `caller` if role gating is done in middleware.

- Validate inputs early — return `ErrValidation`.
- Load the record first, then run ownership/role checks — return `ErrForbidden` on failure.
- Use only the sentinels in `internal/service/errors.go` (`ErrValidation`, `ErrInvalidCredentials`, `ErrUserExists`, `ErrNotFound`, `ErrForbidden`). Never return raw pgx errors; repositories map them to `repository.ErrNotFound` / `repository.ErrConflict`, and services translate those to the service sentinels.

## 3. Repository — `internal/repository/<domain>_repo.go` (only if new SQL is needed)

- Use `r.pool.QueryRow`/`Query`/`Exec`.
- Always map `pgx.ErrNoRows` → `ErrNotFound`.
- On inserts that can conflict on a unique constraint, check `*pgconn.PgError` and map code `23505` → `ErrConflict` (see `UserRepo.Create` for the pattern).
- If the endpoint requires schema changes, also add a new migration file `migrations/00N_<change>.sql`. Migrations only re-run when `pgdata` is wiped — remind the user to run `docker compose down -v && docker compose up --build` locally.

## 4. Interfaces

**Do not skip this** — the tests depend on the fakes in `internal/service/fakes_test.go` implementing these interfaces exactly.

- `internal/service/auth_service.go` (`UserRepo` interface) — add any new user repo methods.
- `internal/service/todo_service.go` (`TodoRepo` interface) — add any new todo repo methods.
- `internal/handler/<x>_handler.go` (`AuthSvc` / `TodoSvc` / `AdminSvc` interface) — add the new service method so the handler stays testable with a stub.

## 5. Handler — same file as the service interface

Standard body shape:
```go
func (h *TodoHandler) DoThing(w http.ResponseWriter, r *http.Request) {
    id, ok := idParam(w, r, "id")   // if a path param
    if !ok { return }
    var body XyzRequest
    if !decodeBody(w, r, &body) { return }   // if a body
    result, err := h.svc.DoThing(r.Context(), id, body.Field, callerFrom(r))
    if err != nil { handleErr(w, err); return }
    writeJSON(w, http.StatusOK, result)
}
```
Use `w.WriteHeader(http.StatusNoContent)` for deletes.

## 6. Swag annotations (above the handler method)

Required tags: `@Summary`, `@Tags`, `@Router <path> [<method>]`, at least one `@Success`, and one `@Failure` per non-2xx status code the handler can return. Protected routes also need `@Security BearerAuth`.

```go
//	@Summary	Do the thing
//	@Tags		lists
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int			true	"List ID"
//	@Param		body	body		XyzRequest	true	"Payload"
//	@Success	200		{object}	StatusResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	403		{object}	ErrorResponse
//	@Failure	404		{object}	ErrorResponse
//	@Security	BearerAuth
//	@Router		/lists/{id}/do-thing [post]
```

## 7. Router — `internal/handler/router.go`

Use the Go 1.22+ mux pattern `"<METHOD> /path/{id}"` and pick the right wrapper:

| Access                              | Wrapper                                                |
| ---                                 | ---                                                    |
| Public                              | `mux.HandleFunc(...)` directly                         |
| Any authenticated user              | `mux.Handle("...", protect(h.Method))`                 |
| power_user + admin                  | `mux.Handle("...", powerRole(h.Method))`               |
| admin only                          | `mux.Handle("...", adminRole(h.Method))`               |

Per-record ownership (owner vs power_user vs admin) is enforced **inside the service**, not by middleware. Middleware `RequireRole` is only for endpoints where the whole route is off-limits to a role.

## 8. Tests

Mirror the service change in `internal/service/*_service_test.go`. Cover:
- Success path.
- Validation error (empty/invalid input).
- Not-found (ID that doesn't exist).
- Forbidden combinations: non-owner as `user`, non-owner as `power_user`, non-owner as `admin` — assert the right role can/can't do the thing per `canRead`/`canWrite` semantics.

If the handler has non-trivial logic beyond delegation, add a stub-based handler test following `internal/handler/auth_handler_test.go`.

## 9. Regenerate Swagger

Invoke the `regenerate-swagger` skill (or run `$(go env GOPATH)/bin/swag init -g main.go -o docs --parseDependency --parseInternal`).

## 10. Verify

```
go vet ./...
go build ./...
go test -race ./...
```

All three must pass before considering the endpoint done.
