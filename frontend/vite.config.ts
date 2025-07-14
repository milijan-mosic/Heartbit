import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 3000,
    host: true,
    allowedHosts: ["app.heartbit.local"],
  },
  plugins: [react(), tailwindcss()],
});
