import React from 'react';
import { Box, Button, Group, PasswordInput, Stack, TextInput } from '@mantine/core';
import { useToast } from '../../../../hooks';
import { useAuthStore } from '../../../../stores';
import { SignUserInDto, useSignUserInMutation } from '../../../../api';
import { hasLength, isEmail, useForm } from '@mantine/form';

type Props = {
    onSubmit: () => void;
    onSwitchToRegister: () => void;
    onSwitchToResetPassword: () => void;
}

export const LoginForm: React.FC<Props> = ({onSubmit, onSwitchToRegister, onSwitchToResetPassword}) => {
    const {successToast, errorToast} = useToast();
    const {login} = useAuthStore();
    const {
        mutate: signIn,
    } = useSignUserInMutation({
        onSuccess: (user) => {
            login(user);
            successToast('Successfully logged in!');
            onSubmit();
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
            password: hasLength({min: 3}, 'Password must consist of at least 3 characters'),
        },
        validateInputOnBlur: true,
    });

    return (
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
                    <Button
                        size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact
                        onClick={onSwitchToRegister}
                    >
                        Create an account
                    </Button>
                    <Button
                        size="sm" classNames={{label: 'text-xs'}} variant="subtle" compact
                        onClick={onSwitchToResetPassword}
                    >
                        Forgot password
                    </Button>
                </Stack>
                <Button type="submit" variant="filled" className="border-violet-600">Log In</Button>
            </Group>
        </Box>
    );
};
