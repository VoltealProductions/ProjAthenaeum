/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/**/*.{html,js,templ,go,lua}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

