"""The sruon.dev family layer: serve STATIC (family.css, glyph.svg, tokens.json, fonts/) and add TEMPLATES to the Jinja loader."""

from pathlib import Path

STATIC = Path(__file__).parent / "static"
TEMPLATES = Path(__file__).parent / "templates"
