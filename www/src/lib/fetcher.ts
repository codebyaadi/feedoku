interface FetchOptions extends RequestInit {
  headers?: Record<string, string>;
}

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export async function fetchWithAuth<T>(
  route: string,
  options: FetchOptions = {}
): Promise<Response> {
  const idxToken = process.env.NEXT_PUBLIC_PROJECT_IDX_TOKEN;

  const url = `${BASE_URL}${route}`;
  console.log('BASE_URL: ', BASE_URL);
  console.log('ROUTE: ', route);
  console.log('URL: ', url);

  const headers: Record<string, string> = {
    ...options.headers,
    'Content-Type': 'application/json',
  };

  if (idxToken) {
    headers.Authorization = `Bearer ${idxToken}`;
  }

  const response = await fetch(url, {
    ...options,
    headers,
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(
      `Error ${response.status}: ${errorText || response.statusText}`
    );
  }

  return response;
}

export const fetcher = <T>(url: string): Promise<T> =>
  fetchWithAuth<T>(url).then((res) => res.json());
