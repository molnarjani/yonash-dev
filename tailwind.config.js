/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/**/*.{html,templ}',
    './node_modules/flowbite/**/*.js',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('flowbite/plugin')
  ],
}
