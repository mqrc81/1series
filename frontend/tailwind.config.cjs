/** @type {import('tailwindcss').Config} */
module.exports = {
    important: true,
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",],
    theme: {
        extend: {
            fontFamily: {
                jost: ['Jost', 'sans-serif'],
                nunito: ['Nunito', 'sans-serif'],
                rubik: ['Rubik', 'sans-serif'],
            },
        },
    },
    plugins: [],
}
