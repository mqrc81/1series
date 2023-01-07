import React from 'react';
import { Avatar, Menu, Modal } from '@mantine/core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowRightFromBracket, faArrowRightToBracket, faCog, faUser } from '@fortawesome/free-solid-svg-icons';
import { useAuthStore } from '../../../stores';
import { useLogoutUserMutation } from '../../../api';
import { useDisclosure } from '@mantine/hooks';
import { useNavigate } from 'react-router-dom';

export const HeaderUserMenu: React.FC = () => {
    const {isLoggedIn, logout} = useAuthStore();
    const {
        mutate: doLogout,
    } = useLogoutUserMutation({
        onSuccess: logout,
    });

    const [loginModalOpened, {open: openLoginModal, close: closeLoginModal}] = useDisclosure(false);

    const navigate = useNavigate();

    return (
        <>
            <Menu position="bottom-end" closeOnItemClick>
                <Menu.Target>
                    <Avatar className="ml-auto mt-3 cursor-pointer bg-violet-600" />
                </Menu.Target>
                <Menu.Dropdown>
                    {isLoggedIn()
                        ? <Menu.Item
                            icon={<FontAwesomeIcon icon={faArrowRightFromBracket} />}
                            onClick={() => doLogout()}
                        >Logout</Menu.Item>
                        : <Menu.Item
                            icon={<FontAwesomeIcon icon={faArrowRightToBracket} />}
                            onClick={openLoginModal}
                        >Login</Menu.Item>}
                    <Menu.Item
                        icon={<FontAwesomeIcon icon={faUser} />}
                        onClick={() => navigate('/profile')}
                    >Profile</Menu.Item>
                    <Menu.Divider />
                    <Menu.Item
                        icon={<FontAwesomeIcon icon={faCog} />}
                        onClick={() => navigate('/profile?tab=preferences')}
                    >Preferences</Menu.Item>
                </Menu.Dropdown>
            </Menu>

            <Modal
                opened={loginModalOpened}
                onClose={closeLoginModal}
                title="Login"
                centered
            >
                Login Modal Content
            </Modal>
        </>
    );
};
