// Package design serves the sruon.dev family layer to Go-served pages: the stylesheet, the glyph, the fonts, and the wordmark.
package design

import (
	"embed"
	"html"
	"html/template"
)

// CSS is family.css. Its font URLs are relative (fonts/<name>.woff2), so serve Fonts beside the stylesheet.
//
//go:embed family.css
var CSS string

// Glyph is the crystal glyph as a standalone SVG, for favicons and images.
//
//go:embed glyph.svg
var Glyph string

// Fonts holds fonts/inter-tight.woff2 and fonts/jetbrains-mono.woff2; http.FileServerFS(Fonts) serves them under /fonts/.
//
//go:embed fonts/*.woff2
var Fonts embed.FS

// Tokens is tokens.json: every CSS token's value, for places CSS cannot reach.
//
//go:embed tokens.json
var Tokens []byte

// Wordmark is the family wordmark, sruon ◆ site, with site linking to the site's root.
func Wordmark(site string) template.HTML {
	return template.HTML(`<span class="wordmark"><a class="wordmark-prefix" href="https://sruon.dev">sruon</a>` +
		`<svg class="wordmark-glyph" width="10" height="10" viewBox="0 0 12 12" aria-hidden="true">` +
		`<path d="M6 0.5 L11.5 6 L6 11.5 L0.5 6 Z" fill="none" stroke="currentColor" stroke-width="1.2"/>` +
		`<path d="M6 3 L9 6 L6 9 L3 6 Z" fill="currentColor" opacity="0.4"/></svg>` +
		`<a class="wordmark-site" href="/">` + html.EscapeString(site) + `</a></span>`)
}
