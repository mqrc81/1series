import React from 'react';
import { NavLink } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
    faArrowRightFromBracket,
    faCalendarDays,
    faCog,
    faFireFlameCurved,
    faUser,
} from '@fortawesome/free-solid-svg-icons';
import { Avatar, Header as AppHeader, Menu } from '@mantine/core';
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
        <AppHeader height={60} className="md:px-60">
            <nav className="grid grid-cols-3 gap-3">
                {/*<NavLink*/}
                {/*    to="/"*/}
                {/*    className="font-semibold mr-8 mt-4"*/}
                {/*>*/}
                {/*    <span className="text-2xl">*/}
                {/*        <span className="text-transparent bg-clip-text bg-gradient-to-b from-violet-600 to-teal-600">NewSeries</span>*/}
                {/*        <span>.top</span>*/}
                {/*    </span>*/}
                {/*</NavLink>*/}
                <div className="mt-5">
                    {navBarItems.map(({title, path, icon}) => (
                        <NavLink
                            key={title}
                            to={path}
                            className={({isActive}) => ('mr-5 font-medium ' + (isActive ? 'text-teal-600' : 'hover:text-violet-600'))}
                        >
                            <FontAwesomeIcon icon={icon}/>
                            <span className="pl-3">{title}</span>
                        </NavLink>
                    ))}
                </div>
                <div className="mt-3"><HeaderSearchBar/></div>
                <Menu position="bottom-end">
                    <Menu.Target>
                        <Avatar className="ml-auto mt-3 cursor-pointer bg-violet-600"/>
                    </Menu.Target>
                    <Menu.Dropdown>
                        <Menu.Item icon={<FontAwesomeIcon icon={faArrowRightFromBracket}/>}>Logout</Menu.Item>
                        <Menu.Item icon={<FontAwesomeIcon icon={faUser}/>}>Account</Menu.Item>
                        <Menu.Divider/>
                        <Menu.Item icon={<FontAwesomeIcon icon={faCog}/>}>Preferences</Menu.Item>
                    </Menu.Dropdown>
                </Menu>
            </nav>
        </AppHeader>
    );
};
