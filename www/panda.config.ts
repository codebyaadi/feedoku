import { defineConfig } from '@pandacss/dev';
import { createPreset } from '@park-ui/panda-preset';

export default defineConfig({
  // Whether to use css reset
  preflight: true,

  // Presets
  presets: [
    '@pandacss/preset-base',
    createPreset({
      accentColor: 'amber',
      grayColor: 'sage',
      borderRadius: 'sm',
    }),
  ],

  // Where to look for your css declarations
  include: ['./src/**/*.{js,jsx,ts,tsx}', './pages/**/*.{js,jsx,ts,tsx}'],

  // Files to exclude
  exclude: [],

  // JS Framework
  jsxFramework: 'react',

  // Useful for theme customization
  theme: {
    extend: {
      tokens: {
        fonts: {
          body: { value: 'var(--font-outfit), sans-serif' },
          code: { value: 'var(--font-roboto-mono), monospace' },
          'geist-sans': { value: 'var(--font-geist-sans), sans-serif' },
          'geist-mono': { value: 'var(--font-geist-mono), monospace' },
        },
      },
    },
  },

  // The output directory for your css system
  outdir: 'styled-system',
});
