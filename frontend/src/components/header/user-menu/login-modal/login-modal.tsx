import React, { useState } from 'react';
import { Box, Button, Group, Modal, PasswordInput, Stack, TextInput } from '@mantine/core';
import { hasLength, isEmail, matches, useForm } from '@mantine/form';
import { SignUserInDto, SignUserUpDto, useSignUserInMutation, useSignUserUpMutation } from '../../../../api';
import { useAuthStore } from '../../../../stores';
import { useToast } from '../../../../hooks';

type Props = { opened: boolean, onClose: () => void }

export const HeaderUserMenuLoginModal: React.FC<Props> = ({opened, onClose}) => {
    const {successToast, errorToast} = useToast();
    const {login} = useAuthStore();
    const {
        mutate: signIn,
    } = useSignUserInMutation({
        onSuccess: (user) => {
            login(user);
            successToast('Successfully logged in!');
            onClose();
        },
        onError: () => errorToast('Invalid credentials!'),
    });
    const {
        mutate: signUp,
    } = useSignUserUpMutation({
        onSuccess: (user) => {
            login(user);
            successToast('Successfully registered!');
            onClose();
        },
        onError: () => errorToast('Invalid credentials!'),
    });

    const [currentForm, setCurrentForm] = useState<'Login' | 'Register' | 'Reset Password'>('Login');

    const loginForm = useForm<SignUserInDto>({
        initialValues: {
            email: '',
            password: '',
        },
        validate: {
            email: isEmail('Invalid email'),
            password: hasLength({min: 3}, 'Password must consist of at least 3 characters'),
        },
        validateInputOnBlur: true,
    });

    const registerForm = useForm<SignUserUpDto>({
        initialValues: {
            username: loginForm.values.email.split('@')[0],
            email: loginForm.values.email,
            password: loginForm.values.password,
        },
        validate: {
            username: matches(/^[a-zA-Z0-9]{3,16}$/, 'Username must consist of 3-16 letters/numbers'),
            email: isEmail('Invalid email'),
            password: hasLength({min: 3}, 'Password must consist of at least 3 characters'),
        },
        validateInputOnBlur: true,
    });

    return (
        <Modal
            opened={opened}
            onClose={onClose}
            title={currentForm}
            centered
        >
            {currentForm === 'Login' && (
                <Box
                    component="form" className="mx-auto max-w-sm"
                    onSubmit={loginForm.onSubmit(values => signIn(values))}
                >
                    <TextInput
                        className="my-3"
                        label="Email" placeholder="example@mail.com" withAsterisk
                        {...loginForm.getInputProps('email')}
                    />
                    <PasswordInput
                        className="my-3"
                        label="Password" placeholder="****" withAsterisk
                        {...loginForm.getInputProps('password')}
                    />

                    <Group position="apart" className="mt-5">
                        <Stack spacing={2} align="start">
                            <Button
                                size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact
                                onClick={() => setCurrentForm('Register')}
                            >
                                Create an account
                            </Button>
                            <Button
                                size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact
                                onClick={() => setCurrentForm('Reset Password')}
                            >
                                Forgot password
                            </Button>
                        </Stack>
                        <Button type="submit" variant="filled" className="border-violet-600">Log In</Button>
                    </Group>
                </Box>
            )}
            {currentForm === 'Register' && (
                <Box
                    component="form" className="mx-auto max-w-sm"
                    onSubmit={registerForm.onSubmit(values => signUp(values))}
                >
                    <TextInput
                        className="my-3"
                        label="Email" placeholder="example@mail.com" withAsterisk
                        {...registerForm.getInputProps('email')}
                    />
                    <TextInput
                        className="my-3"
                        label="Username" placeholder="example" withAsterisk
                        {...registerForm.getInputProps('username')}
                    />
                    <PasswordInput
                        className="my-3"
                        label="Password" placeholder="****" withAsterisk
                        {...registerForm.getInputProps('password')}
                    />

                    <Group position="apart" className="mt-5">
                        <Stack spacing={2} align="start">
                            <Button
                                size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact
                                onClick={() => setCurrentForm('Login')}
                            >
                                I already have an account
                            </Button>
                        </Stack>
                        <Button type="submit" variant="filled" className="border-violet-600">Register</Button>
                    </Group>
                </Box>
            )}
        </Modal>
    );
};
