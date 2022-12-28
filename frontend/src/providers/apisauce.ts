import { ApiResponse, create } from 'apisauce';

export const ApiClient = create({
    baseURL: import.meta.env.VITE_BACKEND_URL + '/api',
    xsrfCookieName: '_csrf',
    xsrfHeaderName: 'X-CSRF-Token',
    withCredentials: true,
});

ApiClient.addResponseTransform((response: ApiResponse<any>) => {
    if (!response.ok) throw response;
});
