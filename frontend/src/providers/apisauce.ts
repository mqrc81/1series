import { ApiResponse, create } from 'apisauce';

export const ApisauceClient = create({
    baseURL: import.meta.env.VITE_BACKEND_URL + '/api',
    xsrfCookieName: '_csrf',
    xsrfHeaderName: 'X-CSRF-Token',
    withCredentials: true,
    timeoutErrorMessage: 'Request timed out',
});

ApisauceClient.addResponseTransform((response: ApiResponse<any>) => {
    if (!response.ok) throw response;
});
