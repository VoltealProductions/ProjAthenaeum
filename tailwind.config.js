/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/views/**/*.{html,js,templ,go,lua}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        'site-archives-img': "url('/public/img/Hallowfall-bg-footer.jpeg')",
        'header-archives-img': "url('/public/img/Hallowfall-bg-header.jpeg')",
        'footer-archives-img': "url('/public/img/Hallowfall-bg.jpeg')",
      },
      backdropBrightness: {
        25: '.25',
        65: '.65',
        70: '.70',
      },
      colors: {
        headerbg: '#cecec9',
        content: '#fcfaea',
        footerbg: '#9e9d93',
        text: '#050909',
        background: '#f0f0f0',
        primary: '#8b70a4',
        secondary: '#c4a2b8',
        accent: '#c2b740',
        accentdarker: '#979363',
        link: '#e092aa',
        visitedlink: '#915366',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

