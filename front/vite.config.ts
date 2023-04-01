import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      jsc: {
        parser: {
          syntax: "ecmascript",
          jsx: true,
        },
      },
    }),
    tsconfigPaths(),
  ],
  /* Https://github.com/vitejs/vite/issues/1037 */
});
