export const getConfig = () => ({
  serverPort: process.env.SERVER_PORT
    ? parseInt(process.env.SERVER_PORT, 10)
    : '',
  clientPort: process.env.VITE_CLIENT_PORT
    ? parseInt(process.env.VITE_CLIENT_PORT, 10)
    : '',
  baseUrl: process.env.VITE_BASE_URL || 'http://localhost:',
});
