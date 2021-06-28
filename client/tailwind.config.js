const colors = require("tailwindcss/colors");

module.exports = {
  // mode: "jit",
  purge: ["./pages/**/*.{js,ts,jsx,tsx}", "./components/**/*.{js,ts,jsx,tsx}"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        ...colors,
      },
      rotate: {
        "-135": "-135deg",
        135: "135deg",
        105: "105deg",
      },
      borderWidth: {
        3: "3px",
      },
    },
  },
  variants: {
    extend: {
      borderWidth: ["hover"],
      opacity: ["disabled"],
    },
  },
  plugins: [require("@tailwindcss/line-clamp")],
};
