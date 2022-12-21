/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",],
    theme: {
        extend: {
            colors: {
                primary: "#5EEAD4",
                secondary: "#EC4899",
                tertiary: "#0bb70c",
                background: "#3F3F3F",
                // background: "#22272E",
            },
        },
    },
    plugins: [],
}
