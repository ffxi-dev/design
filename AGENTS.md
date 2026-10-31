# design

The sruon.dev family layer: tokens, base type, the wordmark, a few shared controls, and the fonts, shipped from one tag as an npm package (lens, res), a Go module (Mogsafe), and a Python package (Magian). `DESIGN.md` is the design document; `README.md` shows how each app installs it. Public, so every app and build fetches it without keys.

## Rules

- The files at the root (`family.css`, `glyph.svg`, `tokens.json`, `fonts/`) are the product. The wrappers (`react/`, `astro/`, `design.go`, `python/`) stay a few lines each and never copy a value the files hold.
- `tokens.json` matches `family.css`'s `:root` exactly. The tests check it.
- Only family rules go here. A rule one site needs stays in that site.
- Releases are tags `vX.Y.Z`, and tags never move. Renaming or removing a token, class, or wrapper is a new major version. Before tagging, set the same version in `package.json` and `pyproject.toml`.
- Fonts stay under the SIL Open Font License, with their licence files beside them.

## Before every commit

```bash
go vet ./... && go test ./...
```

Put `npm` and `uv` on the path so the tests also check the npm and Python packages; CI (`.github/workflows/ci.yml`) always does. After a visible change, open `demo.html` and compare it with lens.

## Part of Vault

One brick of Vault, the owner's platform for capturing retail FFXI and keeping the captures. The platform repository is `ffxi-dev/vault`; its `openspec/changes/add-design-system/` is the working specification. On the owner's PC the repositories sit side by side in `C:/Users/sruon/Documents/GitHub/Vault/`.

- **CI** is written here, not called from vault: a public repository cannot call workflows in vault, which is private.
- **Public.** Never commit credentials or anything that is not meant to be seen.
- **Writing** (docs, commit messages, UI text) follows `ffxi-dev/vault`'s AGENTS.md: short plain words, active voice, no filler.
