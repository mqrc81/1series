import React from 'react';
import { Outlet } from 'react-router-dom';
import { AppShell } from '@mantine/core';
import { Footer, Header } from '../../components';

const Layout: React.FC = () => {
    return (
        <>
            <div className="leading-normal tracking-normal overflow-x-hidden">
                <AppShell hidden>
                    <Header/>
                    <div className="md:px-52 py-10 min-h-[calc(100vh-140px)] w-screen overflow-y-hidden">
                        <Outlet/>
                    </div>
                    <Footer/>
                </AppShell>
            </div>
        </>
    );
};

export default Layout;
