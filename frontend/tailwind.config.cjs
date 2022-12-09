/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",],
    theme: {
        extend: {
            colors: {
                primary: "#3F3F3F",
                secondary: "#5EEAD4",
                tertiary: "#EC4899",
                quaternary: "#0bb70c",
            },
        },
    },
    plugins: [],
}
