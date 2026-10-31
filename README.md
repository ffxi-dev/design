# design

Shared styles for the sruon.dev sites: colours, fonts, buttons, inputs, and the `sruon ◆ site` wordmark. Lens, res, Magian, and Mogsafe all use it. [DESIGN.md](DESIGN.md) has the rules.

Use a tagged version. Renovate opens a pull request when a new one comes out.

## JavaScript

```json
"@ffxi-dev/design": "github:ffxi-dev/design#v1.0.0"
```

```js
import '@ffxi-dev/design/family.css';
import { Wordmark } from '@ffxi-dev/design/react';
import tokens from '@ffxi-dev/design/tokens.json';
```

Import `family.css` before your own CSS. Astro has its own wordmark: `import Wordmark from '@ffxi-dev/design/astro'`.

## Go

```bash
go get github.com/ffxi-dev/design@v1.0.0
```

```go
mux.Handle("GET /fonts/", http.FileServerFS(design.Fonts))
css := design.CSS + siteCSS
wordmark := design.Wordmark("mogsafe")
```

Serve the CSS from the folder that holds `fonts/`: the stylesheet loads `fonts/<name>.woff2` next to itself.

## Python

```toml
[project]
dependencies = ["ffxi-dev-design"]

[tool.uv.sources]
ffxi-dev-design = { git = "https://github.com/ffxi-dev/design", tag = "v1.0.0" }
```

```python
import ffxi_design

app.mount("/design", StaticFiles(directory=ffxi_design.STATIC))
templates = Jinja2Templates(directory=[own_templates, ffxi_design.TEMPLATES])
```

```jinja
{% from "ffxi_design/wordmark.html" import wordmark %}
<link rel="stylesheet" href="/design/family.css">
{{ wordmark("magian") }}
```

## Fonts

Inter Tight and JetBrains Mono, from Fontsource, under the SIL Open Font License (see `fonts/`). Everything else is MIT (see `LICENSE`).

## Testing

```bash
go vet ./... && go test ./...
```

If `npm` and `uv` are installed, the tests also check that the npm and Python packages ship the same `family.css` as the Go module. Open `demo.html` in a browser to see every shared style.
