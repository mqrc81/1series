import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import { NavLink, Outlet } from 'react-router-dom';
import { AppShell, Footer, Header } from '@mantine/core';

const AppLayout: React.FC = () => {
    return (
        <>
            <div className="leading-normal tracking-normal overflow-x-hidden">
                <AppShell>
                    <Header height={60}>
                        <nav className="block flex">
                            <NavLink
                                to="/"
                                className="font-semibold mr-8"
                            >
                                <span className="text-2xl text-transparent bg-clip-text bg-gradient-to-b from-violet-500 to-sky-500">
                                    <span className="">TV Stop</span>
                                </span>
                            </NavLink>
                            <NavLink
                                to="/shows/popular"
                                className={({isActive}) => ('font-semibold mr-8 ' + (isActive ? 'text-sky-500' : 'hover:text-violet-500'))}
                            >
                                <FontAwesomeIcon icon={faFireFlameCurved}/>
                                <span className="pl-3 pb-1">Popular This Week</span>
                            </NavLink>
                            <NavLink
                                to="/shows/releases"
                                className={({isActive}) => ('font-semibold mr-8 ' + (isActive ? 'text-sky-500' : 'hover:text-violet-500'))}
                            >
                                <FontAwesomeIcon icon={faCalendarDays}/>
                                <span className="pl-3">Upcoming Releases</span>
                            </NavLink>
                        </nav>
                    </Header>
                    <div id="content" className="md:px-60 pt-10 h-full w-full overflow-y-hidden">
                        <Outlet/>
                    </div>
                    <Footer height={80}>
                        Zeries ©2022 by Marc Schmidt
                    </Footer>
                </AppShell>
            </div>
        </>
    );
};

export default AppLayout;
