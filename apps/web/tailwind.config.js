/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontSize: {
				h1: ['72px', { fontWeight: 'bold', lineHeight: '1.33em' }],
				h2: ['60px', { lineHeight: '1.33em', fontWeight: 'semibold' }],
				h3: ['54px', { fontWeight: 'bold', lineHeight: '1.6em' }],
				h4: ['36px', { fontWeight: 'bold', lineHeight: '1.6em' }],
				h5: ['34px', { fontWeight: 'bold', lineHeight: '1.5em' }],
				h6: ['30px', { fontWeight: 'bold', lineHeight: '1.5em' }],

				'body-lg': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],
				'body-default': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],
				'body-sm': ['20px', { fontWeight: 'normal', lineHeight: '1.75em' }],

				'caption-1': ['16px', { fontWeight: 'semibold', lineHeight: '1.35em' }],
				'caption-2': ['14px', { fontWeight: 'semibold', lineHeight: '1.4em' }],
				'caption-regular': ['14px', { fontWeight: 'normal', lineHeight: '1.35em' }],

				'btn-md': ['14px', { fontWeight: 'medium', lineHeight: '1.3em' }],
				'btn-lg': ['16px', { fontWeight: 'medium', lineHeight: '1.3em' }],

				'nav-link': ['16px', { fontWeight: 'medium', lineHeight: '1.5em' }]
			},
			fontFamily: {
				dmSans: ['DM Sans', 'system-ui']
			},
			colors: {
				border: 'var(--border)',
				black: 'var(--black-1)',
				'black-2': 'var(--black-2)',
				white: 'var(--white-1)',
				'white-2': 'var(--white-2)',
				primary: 'var(--primary-1)',
				'primary-2': 'var(--primary-2)',
				'primary-3': 'var(--primary-3)',
				'primary-4': 'var(--primary-4)',
				secondary: 'var(--secondary-1)',
				'secondary-2': 'var(--secondary-2)',
				'secondary-3': 'var(--secondary-3)',
				'bg-primary-1': 'var(--bg-primary-1)',
				'bg-primary-2': 'var(--bg-primary-2)',
				'bg-primary-3': 'var(--bg-primary-3)',
				'bg-secondary-1': 'var(--bg-secondary-1)',
				'bg-secondary-2': 'var(--bg-secondary-2)',
				'bg-tertiary-1': 'var(--bg-tertiary-1)',
				'bg-tertiary-2': 'var(--bg-tertiary-2)',
				'content-surface': 'var(--content-surface)',
				'content-primary': 'var(--content-primary)',
				'content-secondary': 'var(--content-secondary)',
				'content-tertiary': 'var(--content-tertiary)'
			}
		}
	},
	plugins: []
};
