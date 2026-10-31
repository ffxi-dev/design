import { createElement as h } from 'react';

// Wordmark is the family wordmark, sruon ◆ site, with site linking to the site's root.
export function Wordmark({ site }) {
  return h(
    'span',
    { className: 'wordmark' },
    h('a', { className: 'wordmark-prefix', href: 'https://sruon.dev' }, 'sruon'),
    h(
      'svg',
      { className: 'wordmark-glyph', width: 10, height: 10, viewBox: '0 0 12 12', 'aria-hidden': true },
      h('path', { d: 'M6 0.5 L11.5 6 L6 11.5 L0.5 6 Z', fill: 'none', stroke: 'currentColor', strokeWidth: 1.2 }),
      h('path', { d: 'M6 3 L9 6 L6 9 L3 6 Z', fill: 'currentColor', opacity: 0.4 }),
    ),
    h('a', { className: 'wordmark-site', href: '/' }, site),
  );
}
