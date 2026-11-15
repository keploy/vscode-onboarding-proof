# The conversion proof's sample project

scripts/convert-proof.js copies this folder into a fresh workspace and drives a
new user's first session through it: the extension's hand-off, then T2
(`keploy mock record|replay` of its tests) and T1 (`keploy record|test` of its
HTTP app).

- `dep/` is the one real dependency: a small HTTP service. The proof starts it
  for each recording and stops it for each replay or test, so only Keploy's
  recording can answer then.
- `shop/` is the client the tests and the app call it through.
- `app/` is the HTTP app T1 records and tests (`GET /user/{id}`,
  `GET /price/{item}`).
- `shop/shop_test.go` is the user's own test, as it is before any agent ran.
- `_agent/` (a folder `go test ./...` skips) is what the coding agent writes
  when the extension hands it the setup: `keploy.yml` and the test, joined to keploy through Go's published
  integration (`keploytest`, keploy/go-sdk). The proof's stand-in agent copies
  it over the workspace when it receives the task.

`go.mod` and `go.sum` already name go-sdk (v3.1.0), so the agent's version
needs no change to them.
