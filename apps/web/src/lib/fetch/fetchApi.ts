import { PUBLIC_API_URL } from '$env/static/public';

type FetchApiResponse<T> = { data?: T & { statusCode: number }; error?: string | string[] | null };

export async function fetchApi<T>(path: string, init?: RequestInit): Promise<FetchApiResponse<T>> {
  let data: FetchApiResponse<T>['data'];
  let error: FetchApiResponse<T>['error'] = null;

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
    if (data && data?.statusCode >= 400) {
      throw data;
    }
  } catch (e) {
    console.error('[FetchAPI]:', e);
    if (e instanceof Error || (e && typeof e === 'object' && 'message' in e)) {
      error = e.message as string | string[];
    }
    return { data: undefined, error };
  }

  return { data, error: undefined };
}
