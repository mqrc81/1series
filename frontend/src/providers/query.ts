import { QueryClient as ReactQueryClient } from 'react-query';

const THIRTY_SECONDS = 30 * 1000;

export const QueryClient = new ReactQueryClient({
    defaultOptions: {
        queries: {
            staleTime: THIRTY_SECONDS,
        },
        mutations: {},
    },
});
