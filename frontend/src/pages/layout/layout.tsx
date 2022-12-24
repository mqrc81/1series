import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import { NavLink, Outlet } from 'react-router-dom';

const AppLayout: React.FC = () => {

    return (
        <>
            <div className="leading-normal tracking-normal overflow-x-hidden">
                <nav className="sticky top-0 z-50 bg-slate-700 w-screen">
                    <div className="md:ml-40 py-5">
                        <NavLink
                            to="/shows/popular"
                            className={({isActive}) => ('font-semibold ' + (isActive ? 'text-secondary' : 'hover:text-primary'))}
                        >
                            <FontAwesomeIcon icon={faFireFlameCurved}/>
                            <span className="pl-3">Popular Right Now</span>
                        </NavLink>
                        <NavLink
                            to="/shows/releases"
                            className={({isActive}) => ('font-semibold ' + (isActive ? 'text-secondary' : 'hover:text-primary'))}
                        >
                            <FontAwesomeIcon icon={faCalendarDays}/>
                            <span className="pl-3">Upcoming Releases</span>
                        </NavLink>
                    </div>
                </nav>
                <div className="md:px-20 bg-background md:mt-16 h-full w-full overflow-y-hidden">
                    <Outlet/>
                </div>
            </div>
        </>
    );
};

export default AppLayout;
