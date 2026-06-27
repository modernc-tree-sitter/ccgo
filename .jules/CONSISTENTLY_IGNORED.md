## IGNORE: Bootstrapping Unrequested CI/CD or Tooling

**- Pattern:** Adding configuration files for GitHub Actions, Renovate, or global tooling (`mise.toml`, etc.) without an explicit user request.
**- Justification:** Agents must execute only the explicitly requested outcome and not expand scope to bundle "nice to have" CI/CD or tooling configurations.
**- Files Affected:** `.github/workflows/*`, `renovate.json`, `mise.toml`

## IGNORE: Unconditional Error Reporting Breaking Verbosity

**- Pattern:** Replacing `_ = err` with unconditional error reporting calls (e.g., `reporter.ReportError(err)`) that omit, bypass, or break existing conditional verbosity logic like `if dmesgs { ... }`.
**- Justification:** Centralized error reporting must strictly preserve existing conditional verbosity logic to avoid regressions that break the tool's expected silent or quiet behavior.
**- Files Affected:** `v4/lib/exec.go`

## IGNORE: Removing Dead Code After `panic(todo(...))`

**- Pattern:** Deleting seemingly unreachable code that follows a `panic(todo(...))` statement or within an exhaustive switch.
**- Justification:** `panic(todo(...))` acts as a placeholder for unimplemented features. The code following it is not true dead code to be cleaned up (e.g., by a Janitor), but rather part of incomplete logic that should remain.
**- Files Affected:** `v3/lib/go.go`, `v4/lib/exec.go`, `v4/lib/expr.go`

## IGNORE: Documenting the "What" Instead of the "Why"

**- Pattern:** Adding obvious godoc comments that simply restate what a function does or its signature, rather than explaining its context, nuance, or rationale.
**- Justification:** Documentation should focus on the "why", the nuances, and the flow of code without adding obvious or redundant comments.
**- Files Affected:** `v4/lib/etc.go`, `v4/lib/exec.go`, `v4/lib/ccgo.go`

## IGNORE: Deletion of `.golden` Testdata Files

**- Pattern:** Deleting or modifying `.golden` testdata files as part of code refactoring or bug fixing.
**- Justification:** Test execution in `v4/lib` can inadvertently overwrite or delete `.golden` testdata files if test outputs or errors change. These unintended modifications must be unstaged before committing.
**- Files Affected:** `v4/lib/testdata/*.golden`
