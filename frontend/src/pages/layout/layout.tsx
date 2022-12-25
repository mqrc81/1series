import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import { NavLink, Outlet } from 'react-router-dom';
import { Layout } from 'antd';
import { Content, Footer, Header } from 'antd/es/layout/layout';

const AppLayout: React.FC = () => {
    return (
        <>
            <div className="leading-normal tracking-normal overflow-x-hidden">
                <Layout className="divide-y divide-white">
                    <Header className="md:px-60">
                        <nav className="block flex">
                            <NavLink
                                to="/"
                                className="font-semibold mr-8"
                            >
                                <span className="text-2xl text-transparent bg-clip-text bg-gradient-to-b from-cyan-400 to-pink-500">
                                    <span className="">ZERIE</span><span
                                    className="fa-flip-horizontal inline-block bg-clip-text bg-gradient-to-b from-cyan-400 to-pink-500"
                                >Z</span>
                                </span>
                            </NavLink>
                            <NavLink
                                to="/shows/popular"
                                className={({isActive}) => ('font-semibold mr-8 ' + (isActive ? 'text-pink-500' : 'hover:text-cyan-400'))}
                            >
                                <FontAwesomeIcon icon={faFireFlameCurved}/>
                                <span className="pl-3 pb-1">Popular This Week</span>
                            </NavLink>
                            <NavLink
                                to="/shows/releases"
                                className={({isActive}) => ('font-semibold mr-8 ' + (isActive ? 'text-pink-500' : 'hover:text-cyan-400'))}
                            >
                                <FontAwesomeIcon icon={faCalendarDays}/>
                                <span className="pl-3">Upcoming Releases</span>
                            </NavLink>
                        </nav>
                    </Header>
                    <Content id="content" className="md:px-60 pt-10 h-full w-full overflow-y-hidden">
                        <Outlet/>
                    </Content>
                    <Footer className="md:px-60 bottom-0 w-full">
                        Zeries ©2022 by Marc Schmidt
                    </Footer>
                </Layout>
            </div>
        </>
    );
};

export default AppLayout;
