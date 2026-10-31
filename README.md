# design

The sruon.dev family layer: one stylesheet (`family.css`), the crystal glyph (`glyph.svg`), the tokens as data (`tokens.json`), the two fonts, and a wordmark for React, Astro, Go, and Jinja. `DESIGN.md` describes the system.

One tag releases it three ways. Pin the tag; Renovate bumps it.

## npm (React, Astro, Vite)

```json
"@ffxi-dev/design": "github:ffxi-dev/design#v1.0.0"
```

```js
import '@ffxi-dev/design/family.css';          // before the site's own CSS; the bundler copies the fonts
import { Wordmark } from '@ffxi-dev/design/react';
import tokens from '@ffxi-dev/design/tokens.json';
```

In Astro: `import Wordmark from '@ffxi-dev/design/astro'`. In CSS: `@import '@ffxi-dev/design/family.css';`.

## Go

```bash
go get github.com/ffxi-dev/design@v1.0.0
```

```go
mux.Handle("GET /fonts/", http.FileServerFS(design.Fonts)) // family.css asks for fonts/<name>.woff2 beside itself
css := design.CSS + siteCSS                                // serve from the folder that holds fonts/, such as /style.css
wordmark := design.Wordmark("mogsafe")                     // template.HTML
```

`design.Glyph` is the SVG and `design.Tokens` the JSON.

## Python (Jinja)

```toml
[project]
dependencies = ["ffxi-dev-design"]

[tool.uv.sources]
ffxi-dev-design = { git = "https://github.com/ffxi-dev/design", tag = "v1.0.0" }
```

```python
import ffxi_design
app.mount("/design", StaticFiles(directory=ffxi_design.STATIC))   # /design/family.css, /design/fonts/...
templates = Jinja2Templates(directory=[own_templates, ffxi_design.TEMPLATES])
```

```jinja
{% from "ffxi_design/wordmark.html" import wordmark %}
<link rel="stylesheet" href="/design/family.css">
{{ wordmark("magian") }}
```

## Fonts

Inter Tight and JetBrains Mono, from Fontsource's variable builds (latin subset), under the SIL Open Font License: `fonts/inter-tight-OFL.txt` and `fonts/jetbrains-mono-OFL.txt`.

## Checks

```bash
go vet ./... && go test ./...
```

With `npm` and `uv` on the path, the tests also pack the npm and Python packages and check that both carry the `family.css` Go embeds. `demo.html` shows every shared class; open it in a browser after a change.
