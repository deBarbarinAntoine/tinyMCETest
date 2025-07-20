/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "../templates/*.{tmpl,html}", // Point to your Go templates
        // If you have other files that use Tailwind classes (e.g., JS files generating HTML),
        // add them here as well.
        // "./public/**/*.html",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        require('daisyui'),
    ],
    daisyui: {
        themes: [
            {
                nord: {
                    ...require("daisyui/src/theming/themes")["nord"],
                    // You can override specific colors for 'nord' here if needed
                },
            },
            {
                sunset: {
                    ...require("daisyui/src/theming/themes")["sunset"],
                    // You can override specific colors for 'sunset' here if needed
                },
            },
        ],
        // You can set the default theme here, but `data-theme` on HTML is more common for dynamic switching.
        // If you want a global default, you might add:
        // darkTheme: "sunset", // For prefers-color-scheme: dark
        // base: true, // adds base styles
        // utils: true, // adds utility classes
        // logs: true, // shows daisyUI logs in console
        // rtl: false, // enables right-to-left (RTL) mode
        // prefix: "", // prefix for daisyUI classes (e.g., 'daisy-btn')
    },
};