import React from 'react';
import { useDisclosure } from '@mantine/hooks';
import { AuthModal } from '../../components';
import { useNavigate } from 'react-router-dom';

const PasswordReset: React.FC = () => {
    const [authModalOpened, {close: closeAuthModal}] = useDisclosure(true);
    const navigate = useNavigate();
    return (
        <>
            <AuthModal
                opened={authModalOpened}
                onClose={() => {
                    closeAuthModal();
                    navigate('/');
                }}
                initialForm="Reset Password"
            />
        </>
    );
};

export default PasswordReset;
