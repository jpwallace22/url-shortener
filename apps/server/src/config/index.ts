export const getConfig = () => ({
  SERVER_PORT: process.env.SERVER_PORT
    ? parseInt(process.env.SERVER_PORT, 10)
    : '',
  CLIENT_PORT: process.env.VITE_CLIENT_PORT
    ? parseInt(process.env.VITE_CLIENT_PORT, 10)
    : '',
  BASE_URL: process.env.VITE_BASE_URL || 'http://localhost:',
  DATABASE_URL: process.env.DATABASE_URL,
});
