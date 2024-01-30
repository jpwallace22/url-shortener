export const getConfig = () => ({
	clientPort: import.meta.env.VITE_CLIENT_PORT
		? parseInt(import.meta.env.VITE_CLIENT_PORT, 10)
		: '',
	baseUrl: import.meta.env.VITE_BASE_URL || 'http://localhost:'
});
