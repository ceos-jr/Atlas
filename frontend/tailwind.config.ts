import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic':
          'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
      colors: {
        'primary-blue': '#2860FF',
        'secondary-blue': '#2196F2',
        'terciary-blue': '#2196F2',
        'primary-orange': '#FF9100',
        'secondary-orange': '#FF8A00',
        'primary-yellow': '#FFC107',
        'primary-gray': '#F4F9FF',
        'secondary-gray': '#727272',
      },
      fontFamily: {
        'quick' : ['Quicksand', 'sans-serif'],
        'roboto' : ['Roboto', 'sans-serif'],
        'roboto-bold' : ['Roboto-bold', 'sans-serif']
      }
    },
  },
  plugins: [],
}
export default config
