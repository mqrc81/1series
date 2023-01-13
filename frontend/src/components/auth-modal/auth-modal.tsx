import React, { useState } from 'react';
import { Modal } from '@mantine/core';
import { AuthModalRegister } from './register/register';
import { AuthModalLogin } from './login/login-form';
import { AuthModalForgotPassword } from './forgot-password/forgot-password';
import { AuthModalResetPassword } from './reset-password/reset-password';

type Props = {
    opened: boolean;
    onClose: () => void;
    initialForm: 'Login' | 'Register' | 'Forgot Password' | 'Reset Password';
}

export const AuthModal: React.FC<Props> = ({opened, onClose, initialForm}) => {

    const [currentForm, setCurrentForm] = useState<Props['initialForm']>(initialForm);

    return (
        <Modal
            opened={opened}
            onClose={onClose}
            title={currentForm}
            centered
        >
            {currentForm === 'Login' && (
                <AuthModalLogin
                    onSubmit={onClose}
                    onSwitchToRegister={() => setCurrentForm('Register')}
                    onSwitchToResetPassword={() => setCurrentForm('Forgot Password')}
                />
            )}
            {currentForm === 'Register' && (
                <AuthModalRegister
                    onSubmit={onClose}
                    onSwitchToLogin={() => setCurrentForm('Login')}
                />
            )}
            {currentForm === 'Forgot Password' && (
                <AuthModalForgotPassword
                    onSubmit={onClose}
                    onSwitchToLogin={() => setCurrentForm('Login')}
                />
            )}
            {currentForm === 'Reset Password' && (
                <AuthModalResetPassword
                    onSubmit={() => setCurrentForm('Login')}
                />
            )}

        </Modal>
    );
};
