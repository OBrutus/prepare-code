# prepare-code
Lightweight CLI helper to prepare competitive programming files from templates.

prepare-code is a small command-line app (written in Go) that:
- detects the online problem platform from a problem URL,
- extracts a short problem code/name,
- creates a new source file from a language-specific template and fills basic metadata (author, date, link, working directory).

**Status (current branch: `port-cli-to-go`)**: Core CLI implemented in Go. Platform detection currently supports Codeforces. Templates are available for Java, C++ and Python in the `template/` folder.

**Quick install**
- Build locally with Python:

```
curl https://raw.githubusercontent.com/OBrutus/prepare-code/refs/heads/mainline/install.py | python3
```

- Or run the provided `build.sh` which wraps the build steps used by this repo.

- There's also a convenience installer `install.py` which copies the repository into `~/.prepare-code` and can add it to your shell `PATH` (supports macOS `zsh` and Linux shells).

**Usage**
- Run the CLI with a problem URL as the first argument:

```
./prepare-code "https://codeforces.com/problemset/problem/1234/A"
```

- If no URL is provided the program prompts for it interactively. The CLI then:
  - identifies the platform (e.g. Codeforces),
  - extracts a short code name (for Codeforces: `contestNumber-A` → `1234-A`),
  - prompts for a language (default: `java`), and
  - creates a source file named `<code>.<lang>` (for example `1234-A.java`) in the current directory populated from `template/<lang>/template.<lang>`.

**Templates**
- Templates live under `template/<lang>/template.<lang>`.
- Current templates included: `java`, `cpp`, `python`.
- Templates use `%s` placeholders that are replaced with:
  1. Author (from local user lookup),
  2. Date (YYYY-MM-DD),
  3. Version (template placeholder),
  4. Link (platform name),
  5. Directory (cwd when run)

**Project structure (high level)**
- `main.go` — CLI entrypoint (parses URL, prompts for language, creates session).
- `src/platform` — platform detection and platform-specific logic (e.g. Codeforces extraction in `codeforces.go`).
- `src/session` — session handling and template processing (creates files using templates).
- `template/` — language templates used to create new problem files.
- `install.py`, `build.sh` — helpers for installation and building.

**Notes & Next steps**
- Only Codeforces URL parsing is implemented today. Adding more platforms (AtCoder, SPOJ, etc.) is straightforward by adding new packages under `src/platform` and wiring them in `src/platform/identifier.go`.
- Contributions welcome: improve URL parsing robustness, add more templates, or add a `--lang` / `--out` CLI flags.

If you'd like, I can (pick one):
- update README with a quick example run and screenshots,
- add a `--help` flag and basic CLI docs in `main.go`, or
- implement AtCoder platform detection next.
