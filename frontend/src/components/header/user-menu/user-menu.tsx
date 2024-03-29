import React from 'react';
import { Avatar, Menu } from '@mantine/core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
    faArrowRightFromBracket,
    faArrowRightToBracket,
    faCog,
    faEye,
    faUser,
} from '@fortawesome/free-solid-svg-icons';
import { useAuthStore } from '../../../stores';
import { useSignUserOutMutation } from '../../../api';
import { useDisclosure } from '@mantine/hooks';
import { useNavigate } from 'react-router-dom';
import { AuthModal } from '../../auth-modal/auth-modal';
import { useToast } from '../../../hooks';

export const HeaderUserMenu: React.FC = () => {
    const {successToast} = useToast();
    const {isLoggedIn, logout} = useAuthStore();
    const {
        mutate: doLogout,
    } = useSignUserOutMutation({
        onSuccess: () => {
            logout();
            successToast('Successfully logged out!');
        },
    });

    const [authModalOpened, {open: openAuthModal, close: closeAuthModal}] = useDisclosure(false);

    const navigate = useNavigate();

    return (
        <>
            <Menu position="bottom-end" closeOnItemClick>
                <Menu.Target>
                    <Avatar className="ml-auto mt-3 cursor-pointer bg-violet-600" />
                </Menu.Target>
                <Menu.Dropdown>
                    {isLoggedIn() ? (
                        <>
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faEye} />}
                                onClick={() => navigate('/users/trackedShows')}
                            >Tracked Series</Menu.Item>
                            <Menu.Divider />
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faUser} />}
                                onClick={() => navigate('/profile')}
                            >Profile</Menu.Item>
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faCog} />}
                                onClick={() => navigate('/profile?tab=preferences')}
                            >Preferences</Menu.Item>
                            <Menu.Divider />
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faArrowRightFromBracket} />}
                                onClick={() => doLogout()}
                            >Sign Out</Menu.Item>
                        </>
                    ) : (
                        <>
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faEye} />}
                                onClick={openAuthModal}
                            >Tracked Series</Menu.Item>
                            <Menu.Divider />
                            <Menu.Item
                                icon={<FontAwesomeIcon icon={faArrowRightToBracket} />}
                                onClick={openAuthModal}
                            >Sign In</Menu.Item>
                        </>
                    )}
                </Menu.Dropdown>
            </Menu>
            <AuthModal
                opened={authModalOpened}
                onClose={closeAuthModal}
                initialForm="Login"
            />
        </>
    );
};
