/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['internal/**/*.{html,templ}'],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
