/** @type {import('tailwindcss').Config} */
export default {
  purge: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    fontFamily: {
      Pretendard: ["Pretendard"],
      D2Coding: ["D2Coding"],
    },
    colors: {
      coco: {
        green_500: "#B9FF47",
        yellow_500: "#FFEF51",
        red_500: "#FF0D64",
        blue_500: "#3AA0FF",
      },
      Navy: {
        50: "#E9E8EA",
        100: "#D2D2D6",
        200: "#BCBBC1",
        300: "#A5A4AD",
        400: "#8F8D98",
        500: "#797784",
        600: "#62606F",
        700: "#4C495B",
        800: "#353346",
        900: "#1F1C32",
      },
    },
    extend: {},
  },
  plugins: [
    require("flowbite/plugin"),
    require("@tailwindcss/typography"),
    require("tailwindcss"),
    require("autoprefixer"),
  ],
};
