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
import { useAuthStore } from './stores';
import { StatusCodes } from 'http-status-codes';

export const Main: React.FC = () => {
    useHandleApiError();
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

ReactDOM.createRoot(document.getElementById('root')).render(<Main />);

const useHandleApiError = () => {
    const {errorToast} = useToast();
    const {logout} = useAuthStore();
    useEffect(() => {
        const onError = (error: unknown) => {
            if (isClientError(error)) {
                errorToast(error.data.errorMessage);
                if (error.status === StatusCodes.UNAUTHORIZED) {
                    logout();
                }
            } else {
                errorToast('An unknown error occurred.');
            }
        };
        QueryClient.getDefaultOptions().queries.onError = onError;
        QueryClient.getDefaultOptions().mutations.onError = onError;
    });
};

const isClientError = (error: unknown): error is ApiErrorResponse<ClientError> => {
    return !!((error as ApiErrorResponse<ClientError>)?.data?.errorMessage);
};
