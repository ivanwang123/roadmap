const colors = require("tailwindcss/colors");

module.exports = {
  // mode: "jit",
  purge: ["./pages/**/*.{js,ts,jsx,tsx}", "./components/**/*.{js,ts,jsx,tsx}"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        ...colors,
        secondary: "#F8F8F8", // "#F5F7FB",
        tertiary: "#FAFAFA", // "#F6F8FA",
        hover: colors.blue["400"],
      },
      rotate: {
        "-135": "-135deg",
        135: "135deg",
        105: "105deg",
      },
      borderWidth: {
        3: "3px",
      },
      boxShadow: {
        light:
          "0 0px 8px -5px rgba(0, 0, 0, 0.08), 0 3px 20px -5px rgba(0, 0, 0, 0.06)",
      },
    },
  },
  variants: {
    extend: {
      borderWidth: ["hover"],
      opacity: ["disabled"],
      fontWeight: ["hover"],
      backgroundColor: ["hover", "active"],
    },
  },
  plugins: [require("@tailwindcss/line-clamp")],
};
