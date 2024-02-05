import { PUBLIC_API_URL } from '$env/static/public';

export async function fetchApi<T>(
  path: string,
  init?: RequestInit
): Promise<{ data: T; loading: boolean; error?: string }> {
  let data = null;
  let loading = true;
  let error;

  const url = new URL(path, PUBLIC_API_URL);

  try {
    const res = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        ...init?.headers
      },
      ...init
    });
    data = await res.json();
  } catch (e) {
    console.error('[FetchAPI]:', error);
    if (e instanceof Error) {
      if (Array.isArray(e.message)) {
        error = e.message.join(', ');
      }
      error = e.message || 'An error occurred';
    }
  } finally {
    loading = false;
  }

  return { data, loading, error };
}
