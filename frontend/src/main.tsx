import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { QueryClientProvider } from 'react-query';
import { QueryClient } from './api';
import { ConfigProvider as AntConfigProvider } from 'antd';
import { AntTheme } from './theme';
import { RouterProvider } from 'react-router-dom';
import { Router } from './routing';

document.body.classList.add('bg-background');
ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <QueryClientProvider client={QueryClient}>
            <AntConfigProvider theme={AntTheme}>
                <RouterProvider router={Router}/>
            </AntConfigProvider>
        </QueryClientProvider>
    </React.StrictMode>,
);
