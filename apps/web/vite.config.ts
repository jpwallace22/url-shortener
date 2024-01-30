import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

const base = process.env.VITE_BASE_URL || '';
const server = process.env.SERVER_PORT || '';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	build: {
		commonjsOptions: {
			include: [/@repo-ui/, /node_modules/]
		}
	},
	server: {
		proxy: {
			'/api': `${base}${server}`
		}
	}
});
