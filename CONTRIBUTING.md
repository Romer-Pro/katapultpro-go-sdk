# Contributing to katapultpro-go-sdk

Thank you for considering contributing. This project uses a pull-request workflow: all changes are submitted via PRs, must pass CI, and require maintainer approval before merging.

## How to contribute

1. **Open an issue** (optional but helpful for bugs and features)
   - Bugs: use the [Bug report](.github/ISSUE_TEMPLATE/bug_report.md) template.
   - Features: use the [Feature request](.github/ISSUE_TEMPLATE/feature_request.md) template.

2. **Fork the repo** and create a branch from `main` for your change.

3. **Make your changes**
   - Follow existing style (format with `gofmt` / `just fmt`).
   - Add or update tests as needed.
   - Keep commits focused and messages clear.

4. **Run checks locally**
   - From the repo root (with [just](https://github.com/casey/just) installed):  
     `just ci`  
     This runs `go mod tidy`, build, `go vet`, and tests with the race detector.
   - Without just:
     ```bash
     cd v3 && go mod tidy && cd ..
     go build ./v3/...
     go vet ./v3/...
     go test -race ./v3/...
     ```

5. **Open a pull request** into `main`
   - Describe what you changed and why.
   - Link any related issue (e.g. “Fixes #123”).
   - Ensure the PR title and description are clear for maintainers and history.

## Pull request requirements

- **All CI checks must pass.** The [CI](.github/workflows/ci.yml) workflow runs on every push to your PR (build, vet, tests with race detector). The PR cannot be merged until it passes.
- **Review and approval.** A project maintainer must review and approve the PR before it can be merged into `main`.
- **No direct pushes to `main`.** All changes must go through a PR. If you have write access, still use a branch and open a PR.

## Branch protection (main)

The repository is configured so that:

- Changes to `main` must be made via pull requests.
- The **CI** workflow must succeed before a PR can be merged.
- At least one approval from a project maintainer is required to merge.

If you are a maintainer and need to set or update these rules, go to **Settings → Branches → Branch protection rules** for `main` and require:

- “Require a pull request before merging” (e.g. 1 approval).
- “Require status checks to pass before merging” and select the **Test** (or **CI**) check.
- “Do not allow bypassing the above settings” if you want to enforce this for everyone.

## Development setup

- **Go:** Version is defined in `v3/go.mod` (e.g. 1.25.4). Use that or a compatible version.
- **Tasks:** Optional [just](https://github.com/casey/just) recipes are in the root `justfile` (`just`, `just test`, `just ci`, etc.).
- **Layout:** SDK code lives under `v3/`; design and user docs are in `docs/`. See [docs/design.md](docs/design.md) for versioning and layout.

## Code and conduct

- Be respectful and constructive in issues and PRs.
- By contributing, you agree that your contributions will be licensed under the same license as the project (see [LICENSE](LICENSE)).

## Questions

If something is unclear, open an issue with the “question” label or describe it in a feature-request issue.
