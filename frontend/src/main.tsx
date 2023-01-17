import React, { useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { QueryClientProvider } from 'react-query';
import { RouterProvider } from 'react-router-dom';
import { QueryClient } from './providers/query';
import { Router } from './providers/router';
import { MantineProvider } from '@mantine/core';
import { MantineTheme } from './providers/theme';
import { NotificationsProvider } from '@mantine/notifications';
import { useToast } from './hooks';
import { ApiErrorResponse } from 'apisauce';
import { ClientError } from './api';

export const Main: React.FC = () => {
    useToastOnApiError();
    return (
        <React.StrictMode>
            <QueryClientProvider client={QueryClient}>
                <MantineProvider withGlobalStyles withNormalizeCSS theme={MantineTheme}>
                    <NotificationsProvider limit={3} position="top-right">
                        <RouterProvider router={Router} />
                    </NotificationsProvider>
                </MantineProvider>
            </QueryClientProvider>
        </React.StrictMode>
    );
};
ReactDOM.createRoot(document.getElementById('root')).render(
    <Main />,
);

const useToastOnApiError = () => {
    const {errorToast} = useToast();
    useEffect(() => {
        const onErrorFn = (error: unknown) => {
            errorToast((error as ApiErrorResponse<ClientError>).data?.errorMessage ?? 'An unknown error occurred.');
        };
        QueryClient.getDefaultOptions().queries.onError = onErrorFn;
        QueryClient.getDefaultOptions().mutations.onError = onErrorFn;
    });
};
