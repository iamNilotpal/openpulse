import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
  plugins: [sveltekit()],
  build: {
    commonjsOptions: {
      include: [/@openpulse\/ui/, /node_modules/],
    },
  },
});
