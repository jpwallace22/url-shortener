export const load = async () => {
	const res = await fetch('http://localhost:5173/api');
	return res.json();
};
