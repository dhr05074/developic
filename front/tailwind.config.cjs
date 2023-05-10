/** @type {import('tailwindcss').Config} */
export default {
    purge: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        fontFamily: {
            Pretendard: ["Pretendard"],
        },
        colors: {
            coco: {
                green_500: "#B9FF47",
                yellow_500: "#FFEF51",
                red_500: "#FF0D64",
                blue_500: "3AA0FF",
            },
            Navy: {
                50: "#E6E5EC",
                100: "#CECCD9",
                200: "#B5B2C7",
                300: "#9D99B4",
                400: "#8480A1",
                500: "#6B668E",
                600: "#534C7B",
                700: "#3A3369",
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
