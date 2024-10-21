import { fontFamily } from 'tailwindcss/defaultTheme';

/** @type {import('tailwindcss').Config} */
const config = {
  safelist: ['dark'],
  darkMode: ['class'],
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px',
      },
    },
    extend: {
      colors: {
        red: 'hsl(var(--red) / <alpha-value>)',
        blue: 'hsl(var(--blue) / <alpha-value>)',
        green: 'hsl(var(--green) / <alpha-value>)',
        yellow: 'hsl(var(--yellow) / <alpha-value>)',
        ring: 'hsl(var(--ring) / <alpha-value>)',
        input: 'hsl(var(--input) / <alpha-value>)',
        border: 'hsl(var(--border) / <alpha-value>)',
        grey: {
          50: 'hsl(var(--grey-50) / <alpha-value>)',
          100: 'hsl(var(--grey-100) / <alpha-value>)',
          200: 'hsl(var(--grey-200) / <alpha-value>)',
          300: 'hsl(var(--grey-300) / <alpha-value>)',
          400: 'hsl(var(--grey-400) / <alpha-value>)',
          500: 'hsl(var(--grey-500) / <alpha-value>)',
          600: 'hsl(var(--grey-600) / <alpha-value>)',
          700: 'hsl(var(--grey-700) / <alpha-value>)',
          800: 'hsl(var(--grey-800) / <alpha-value>)',
          900: 'hsl(var(--grey-900) / <alpha-value>)',
        },
        background: {
          DEFAULT: 'hsl(var(--background) / <alpha-value>)',
          100: 'hsl(var(--bg-100) / <alpha-value>)',
          200: 'hsl(var(--bg-200) / <alpha-value>)',
          300: 'hsl(var(--bg-300) / <alpha-value>)',
          400: 'hsl(var(--bg-400) / <alpha-value>)',
          500: 'hsl(var(--bg-500) / <alpha-value>)',
          reverse: 'hsl(var(--reverse-background) / <alpha-value>)',
        },
        foreground: {
          DEFAULT: 'hsl(var(--foreground) / <alpha-value>)',
          200: 'hsl(var(--foreground-200) / <alpha-value>)',
          300: 'hsl(var(--foreground-300) / <alpha-value>)',
          400: 'hsl(var(--foreground-400) / <alpha-value>)',
        },
        primary: {
          DEFAULT: 'hsl(var(--primary) / <alpha-value>)',
          foreground: 'hsl(var(--primary-foreground) / <alpha-value>)',
        },
        secondary: {
          DEFAULT: 'hsl(var(--secondary) / <alpha-value>)',
          foreground: 'hsl(var(--secondary-foreground) / <alpha-value>)',
        },
        destructive: {
          DEFAULT: 'hsl(var(--destructive) / <alpha-value>)',
          foreground: 'hsl(var(--destructive-foreground) / <alpha-value>)',
        },
        muted: {
          DEFAULT: 'hsl(var(--muted) / <alpha-value>)',
          foreground: 'hsl(var(--muted-foreground) / <alpha-value>)',
        },
        accent: {
          DEFAULT: 'hsl(var(--accent) / <alpha-value>)',
          foreground: 'hsl(var(--accent-foreground) / <alpha-value>)',
        },
        popover: {
          DEFAULT: 'hsl(var(--popover) / <alpha-value>)',
          foreground: 'hsl(var(--popover-foreground) / <alpha-value>)',
        },
        card: {
          DEFAULT: 'hsl(var(--card) / <alpha-value>)',
          foreground: 'hsl(var(--card-foreground) / <alpha-value>)',
        },
      },
      borderRadius: {
        sm: 'var(--rounding-sm)',
        md: 'var(--rounding-md)',
        lg: 'var(--rounding-lg)',
        xs: 'var(--rounding-xs)',
        xl: 'var(--rounding-xl)',
        xxll: 'var(--rounding-xxll)',
        full: 'var(--rounding-full)',
        none: 'var(--rounding-none)',
      },
      fontSize: {
        h1: ['72px', { fontWeight: 'bold', lineHeight: '1.33em' }],
        h2: ['60px', { lineHeight: '1.33em', fontWeight: 'semibold' }],
        h3: ['54px', { fontWeight: 'bold', lineHeight: '1.6em' }],
        h4: ['36px', { fontWeight: 'bold', lineHeight: '1.6em' }],
        h5: ['34px', { fontWeight: 'bold', lineHeight: '1.5em' }],
        h6: ['30px', { fontWeight: 'bold', lineHeight: '1.5em' }],

        'body-lg': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],
        'body-sm': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],
        'body-default': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],

        'caption-1': ['16px', { fontWeight: 'semibold', lineHeight: '1.35em' }],
        'caption-2': ['14px', { fontWeight: 'semibold', lineHeight: '1.4em' }],
        'caption-regular': ['14px', { fontWeight: 'normal', lineHeight: '1.35em' }],

        'btn-md': ['14px', { fontWeight: 'medium', lineHeight: '1.3em' }],
        'btn-lg': ['16px', { fontWeight: 'medium', lineHeight: '1.3em' }],

        anchor: ['16px', { fontWeight: 'medium', lineHeight: '1.5em' }],
      },
      fontFamily: {
        sans: ['DM Sans', 'system-ui', ...fontFamily.sans],
      },
      animation: {
        shine: 'shine 2s linear infinite',
        'spin-around': 'spin-around calc(var(--speed) * 2) infinite linear',
        magicslide: 'magicslide var(--speed) ease-in-out infinite alternate',
      },
      keyframes: {
        shine: {
          from: { backgroundPosition: '0 0' },
          to: { backgroundPosition: '-200% 0' },
        },
        'spin-around': {
          '0%': {
            transform: 'translateZ(0) rotate(0)',
          },
          '15%, 35%': {
            transform: 'translateZ(0) rotate(90deg)',
          },
          '65%, 85%': {
            transform: 'translateZ(0) rotate(270deg)',
          },
          '100%': {
            transform: 'translateZ(0) rotate(360deg)',
          },
        },
        magicslide: {
          to: { transform: 'translate(calc(100cqw - 100%), 0)' },
        },
      },
    },
  },
};

export default config;
