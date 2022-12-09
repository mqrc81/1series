import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './app';
import './index.css';
import { QueryClientProvider } from 'react-query';
import { QueryClient } from './api';
import { ConfigProvider as AntConfigProvider } from 'antd';
import { AntTheme } from './theme';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <QueryClientProvider client={QueryClient}>
            <AntConfigProvider theme={AntTheme}>
                <App/>
            </AntConfigProvider>
        </QueryClientProvider>
    </React.StrictMode>,
);
