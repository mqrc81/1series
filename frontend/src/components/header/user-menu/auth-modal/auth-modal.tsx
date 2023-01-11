import React, { useState } from 'react';
import { Modal } from '@mantine/core';
import { RegisterForm } from './register-form';
import { LoginForm } from './login-form';
import { ResetPasswordForm } from './reset-password-form';

type Props = { opened: boolean, onClose: () => void }

export const HeaderUserMenuAuthModal: React.FC<Props> = ({opened, onClose}) => {

    const [currentForm, setCurrentForm] = useState<'Login' | 'Register' | 'Reset Password'>('Login');

    return (
        <Modal
            opened={opened}
            onClose={onClose}
            title={currentForm}
            centered
        >
            {currentForm === 'Login' && (
                <LoginForm
                    onSubmit={onClose}
                    onSwitchToRegister={() => setCurrentForm('Register')}
                    onSwitchToResetPassword={() => setCurrentForm('Reset Password')}
                />
            )}
            {currentForm === 'Register' && (
                <RegisterForm
                    onSubmit={onClose}
                    onSwitchToLogin={() => setCurrentForm('Login')}
                />
            )}
            {currentForm === 'Reset Password' && (
                <ResetPasswordForm
                    onSubmit={() => {}}
                    onSwitchToLogin={() => setCurrentForm('Login')}
                    onSwitchToRegister={() => setCurrentForm('Register')}
                />
            )}
        </Modal>
    );
};
