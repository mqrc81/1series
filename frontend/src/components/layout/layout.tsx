import React from 'react';
import { Content, Header } from 'antd/es/layout/layout';
import { Layout, Menu, MenuProps } from 'antd';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarDays, faFireFlameCurved } from '@fortawesome/free-solid-svg-icons';
import PopularShows from '../popular-shows/popular-shows';

const MenuItems: MenuProps['items'] = [
    {
        key: '1',
        title: 'Popular',
        icon: <FontAwesomeIcon icon={faFireFlameCurved}/>,
    },
    {
        key: '2',
        title: 'Upcoming Releases',
        icon: <FontAwesomeIcon icon={faCalendarDays}/>,
    },
];

const AppLayout: React.FC = () => {
    return (
        <>
            <Layout>
                <Header>
                    <Menu theme="dark" mode="horizontal" items={MenuItems} defaultSelectedKeys={['1']}/>
                </Header>
                <Content>
                    <PopularShows/>
                </Content>
            </Layout>
        </>
    );
};

export default AppLayout;
