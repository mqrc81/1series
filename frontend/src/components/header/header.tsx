import React from 'react';
import { NavLink } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import { Header as AppHeader } from '@mantine/core';
import { HeaderSearchBar } from './search-bar/search-bar';
import { IconProp } from '@fortawesome/fontawesome-svg-core';

type NavBarItem = { title: string, path: string, icon: IconProp }

const navBarItems: NavBarItem[] = [
    {

        title: 'Popular This Week',
        path: '/shows/popular',
        icon: faFireFlameCurved,
    },
    {
        title: 'Upcoming Releases',
        path: '/shows/releases',
        icon: faCalendarDays,
    },
];

export const Header: React.FC = () => {
    return (
        <AppHeader height={60} className="md:px-52">
            <nav className="block flex">
                <NavLink
                    to="/"
                    className="font-semibold mr-8 mt-4"
                >
                    <span className="text-2xl">
                        <span className="text-transparent bg-clip-text bg-gradient-to-b from-violet-500 to-sky-500">NewSeries</span>
                        <span>.top</span>
                    </span>
                </NavLink>
                {navBarItems.map(({title, path, icon}) => (
                    <NavLink
                        key={title}
                        to={path}
                        className={({isActive}) => ('font-semibold mr-8 mt-5 ' + (isActive ? 'text-sky-500' : 'hover:text-violet-500'))}
                    >
                        <FontAwesomeIcon icon={icon}/>
                        <span className="pl-3">{title}</span>
                    </NavLink>
                ))}
                <div className="mt-3"><HeaderSearchBar/></div>
            </nav>
        </AppHeader>
    );
};
