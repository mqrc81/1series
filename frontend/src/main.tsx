import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { QueryClientProvider } from 'react-query';
import { ConfigProvider as AntConfigProvider } from 'antd';
import { RouterProvider } from 'react-router-dom';
import { AntTheme, QueryClient, Router } from './providers';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <QueryClientProvider client={QueryClient}>
            <AntConfigProvider theme={AntTheme}>
                <RouterProvider router={Router}/>
            </AntConfigProvider>
        </QueryClientProvider>
    </React.StrictMode>,
);
