# sruon.dev design

The sruon.dev sites (lens, res, Magian, Mogsafe) are one family. They share tokens, type, the wordmark, and a few controls, all in `family.css`. Each site picks its own density and components on top.

The system was first drawn as HTML, CSS, and JSX prototypes outside these repositories; lens shipped it first, and this document moved here from `lens/docs/design-system.md` when the family layer became its own package.

## Theme

One dark theme, One Dark's canonical `#21252b` base. A lighter base was tried and read washed out. `family.css` sets `color-scheme: dark`, so pages stay dark when the OS asks for light. A light theme would be a minor release.

## Tokens

Every colour and font is a CSS custom property on `:root`. Components read the variables and never hard-code a hex value. `tokens.json` holds the same values for places CSS cannot reach, such as OG cards drawn with Satori; a test keeps the two equal.

```
--bg-base:    #21252b   page background
--bg-surface: #282c34   panels, toolbars, status bars
--bg-raised:  #2c313a   buttons, raised cards
--bg-sunken:  #1b1f25   input wells, data grids, state-frame icons
--bg-hover:   #323842   row hover

--border:     #3a4049   hairline 1px borders
--border-s:   #4b5263   stronger borders (selected, focused, buttons)

--fg:         #abb2bf   main text
--fg-2:       #828997   secondary
--fg-3:       #5c6370   labels
--fg-muted:   #4b5263   placeholders, disabled

--accent:      #61afef  the one accent: links, focus, active rows, key values
--accent-glow: rgba(97, 175, 239, 0.35)
--accent-dim:  #3a6a90  focus ring on inputs
--gold:        #e5c07b  highlighted numbers inside data views only

--ok:   #98c379   success, added
--warn: #e5c07b   attention, changed
--err:  #e06c75   errors, removed, disconnected

--sel:        rgba(97, 175, 239, 0.1)    selected row
--sel-strong: rgba(97, 175, 239, 0.18)   text selection

--font-mono:  'JetBrains Mono', ...
--font-sans:  'Inter Tight', ...
```

Sites may add their own names on top, as aliases or for meanings the family does not have. lens keeps `--in` and `--out` for packet direction (`var(--ok)` and `var(--warn)`); res keeps `--diff-chg` and `--gold-dim`.

## Fonts

Inter Tight (sans) and JetBrains Mono (mono), both under the SIL Open Font License (`fonts/*-OFL.txt`). The package ships each as one variable woff2 file, latin subset, and `family.css` loads them from `./fonts/`. Every site serves them itself: no request to Google. Never use Inter, Roboto, or system-ui.

## Type roles

Each site sets two roles in its own stylesheet and uses them everywhere:

- `--<site>-prose`: page chrome (toolbar text, tabs, status bar, button labels, empty-state text, column headers).
- `--<site>-data`: technical content (ids, hex bytes, offsets, field values, timestamps).

| Site | prose | data |
|---|---|---|
| lens | mono | mono |
| res | sans | mono |
| Magian | sans | mono |
| Mogsafe | sans | mono |

Whatever the roles say, pills (`.chip`), labels (`.label`), and `<kbd>` are always mono, and the wordmark is always sans.

## The wordmark

`sruon ◆ <site>` in the page chrome, top left. `sruon` links to `https://sruon.dev`; the site name links to the site's root. The ◆ is the crystal glyph (`glyph.svg`) in `--accent`. Use the package's wrapper rather than writing the markup:

- React: `import { Wordmark } from '@ffxi-dev/design/react'`, then `<Wordmark site="lens" />`.
- Astro: `import Wordmark from '@ffxi-dev/design/astro'`, then `<Wordmark site="res" />`.
- Go: `design.Wordmark("mogsafe")` returns `template.HTML`.
- Jinja: `{% from "ffxi_design/wordmark.html" import wordmark %}`, then `{{ wordmark("magian") }}`.

## Shared components

| Class or element | Use |
|---|---|
| `.wordmark` | The family wordmark. |
| `.toolbar` | The bar across the top of a page: 36px, `--bg-surface`, a hairline below, items 12px apart. |
| `button`, `button.primary`, `button.ghost` | Squared, 26px, 11.5px mono, 3px radius. |
| `input` (text-like types) | 22px, 11.5px mono, 2px radius, sunken well, `--accent-dim` focus ring. |
| `.label` | Column headers (put it on `th`) and field names: 9.5px uppercase mono, 0.06em spacing, `--fg-3`. |
| `.chip` with `.green`, `.amber`, `.blue`, `.red` | Pills, 18px, 10.5px mono. Plain `.chip` is slate. |
| `<kbd>` | Keyboard hint: 16px, 10px mono, raised background. |
| `.state-frame` | Empty, error, and 404 panel: a 56px icon in corner brackets, a mono title, a body capped at 380px, and an action row. |

A site restyles any of these for its own density in its own stylesheet, loaded after `family.css`.

The element rules (`button`, `input`, `kbd`) style every such element on the page. A site that sizes its own controls on a CSS reset, as res does on Tailwind, imports `family.css` into a cascade layer after the reset and reverts those three element rules inside that layer; res's `src/styles/global.css` shows how.

### State frames

Markup: `.state-frame` holds `.state-frame-icon` (four `.state-bracket` spans, `tl`, `tr`, `bl`, `br`, then the icon glyph), `.state-frame-title`, `.state-frame-body`, and `.state-frame-action`. Use a glyph as the icon: `◯` waiting, `∅` nothing matches, `⚠` error (set the icon's colour to `--err`).

## Rules

- **One accent.** A new meaning uses an existing pill tone, not a new colour.
- **No gradients, glass, or glows.** The one glow is `text-shadow: 0 0 6px var(--accent-glow)` on an active row cursor.
- **No emoji as UI.** Use unicode glyphs at mono size: `◆ ◯ ◇ ▶ ▾ ↗ ⏵ ⚠ ∅`.
- **Square corners.** Nothing rounder than 4px.
- **`--gold` only for data.** Highlighted numbers inside hex or field views, never chrome.
- **No zebra stripes.** Hairline 1px borders only. A tinted row must mean something.
- **Data is mono.** Never set data in body prose sizes.
- **Corner brackets only in state frames.**

## Sites

### lens

All mono, dense, a developer tool.

- Body 12.5px, line height 1.45. Table rows 22px, 11.5px text. Column headers 9.5px. Status bar 22px, 10.5px text. Inputs 22px.
- Its own components: `.statusbar` (22px), `.session-tabs`, `.filter-bar`, `.search-bar`, `.packet-row` with `.row-cursor`, `.field-row`, `.hex-grid` with `.hex-header`, `.byte`, and `.ac`, `.timeline`, `.watch-bar`, and `.banner`.
- Its state frames: `◯` waiting for packets, `∅` no packets match these filters, `∅` no matches for a search, and `⚠` (in `--err`) when schemas fail to load.

### res

Sans chrome, mono data, AG Grid tables, 13.5px body.

### Magian and Mogsafe

Sans chrome, mono data. Magian's request pages bracket their main panel, the one exception to the bracket rule. Mogsafe uses lens's density.
