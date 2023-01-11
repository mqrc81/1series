import React from 'react';
import { Box, Button, Group, Modal, PasswordInput, Stack, TextInput } from '@mantine/core';
import { hasLength, isEmail, useForm } from '@mantine/form';
import { SignUserInDto, useSignUserInMutation } from '../../../../api';
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
        },
        onError: () => errorToast('Invalid credentials!'),
    });

    const form = useForm<SignUserInDto>({
        initialValues: {
            email: '',
            password: '',
        },
        validate: {
            email: isEmail('Invalid email'),
            password: hasLength({min: 2}, 'Invalid password'),
        },
        validateInputOnBlur: true,
    });

    return (
        <Modal
            opened={opened}
            onClose={onClose}
            title="Login"
            centered
        >
            <Box
                component="form" className="mx-auto max-w-sm"
                onSubmit={form.onSubmit(values => signIn(values))}
            >
                <TextInput
                    className="my-3"
                    label="Email" placeholder="example@mail.com" withAsterisk
                    {...form.getInputProps('email')}
                />
                <PasswordInput
                    className="my-3"
                    label="Password" placeholder="****" withAsterisk
                    {...form.getInputProps('password')}
                />

                <Group position="apart" className="mt-5">
                    <Stack spacing={2} align="start">
                        <Button size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact>
                            Create an account
                        </Button>
                        <Button size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact>
                            Forgot password
                        </Button>
                    </Stack>
                    <Button type="submit" variant="filled" className="border-violet-600">Log In</Button>
                </Group>
            </Box>
        </Modal>
    );
};
