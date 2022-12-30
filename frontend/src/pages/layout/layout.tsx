import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import { NavLink, Outlet } from 'react-router-dom';
import { AppShell, Footer, Header } from '@mantine/core';

const AppLayout: React.FC = () => {
    return (
        <>
            <div className="leading-normal tracking-normal overflow-x-hidden">
                <AppShell hidden>
                    <Header height={60} className="md:px-60">
                        <nav className="block flex">
                            <NavLink
                                to="/"
                                className="font-semibold mr-8"
                            >
                                <span className="text-2xl text-transparent bg-clip-text bg-gradient-to-b from-violet-500 to-sky-500">
                                    <span className="">NewSeries.top</span>
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
                    <div className="md:px-60 pt-10 w-full overflow-y-hidden flex-grow">
                        <Outlet/>
                    </div>
                    <Footer height={80} className="md:px-60">
                        TODO Footer
                    </Footer>
                </AppShell>
            </div>
        </>
    );
};

export default AppLayout;
