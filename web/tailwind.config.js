/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        'node_modules/daisyui/dist/**/*.js',
        'node_modules/react-daisyui/dist/**/*.js',
        './src/components/**/*.{js,ts,jsx,tsx,mdx}',
        './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
        './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    ],
    daisyui: {
        themes: ['light', 'dark', 'black'],
    },
    theme: {
        extend: {
            backgroundImage: {
                'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
                'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
            },
        },
    },
    plugins: [require('@tailwindcss/typography'), require('daisyui')],
};
