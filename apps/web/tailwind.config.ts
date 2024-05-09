import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
        "hero-pattern": "/konan-hero.webp",
      },
      colors: {
        xuedao_pink: "#FFFAF4",
        joinus_btn: "#FA9382",
      },
      spacing: {
        '120': '28rem',
      }
    },
  },
  plugins: [],
};
export default config;
