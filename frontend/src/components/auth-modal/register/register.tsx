import React from 'react';
import { useToast } from '../../../hooks';
import { useAuthStore } from '../../../stores';
import { SignUserUpDto, useSignUserUpMutation } from '../../../api';
import { hasLength, isEmail, matches, useForm } from '@mantine/form';
import { Box, Button, Group, PasswordInput, TextInput } from '@mantine/core';

type Props = {
    onSubmit: () => void;
    onSwitchToLogin: () => void;
}

export const AuthModalRegister: React.FC<Props> = ({onSubmit, onSwitchToLogin}) => {
    const {successToast} = useToast();
    const {login} = useAuthStore();
    const {
        mutate: signUp,
        isLoading,
    } = useSignUserUpMutation({
        onSuccess: (user) => {
            login(user);
            successToast('Successfully registered!');
            onSubmit();
        },
    });

    const form = useForm<SignUserUpDto>({
        initialValues: {
            username: '',
            email: '',
            password: '',
        },
        validate: {
            username: matches(/^[a-zA-Z0-9]{3,16}$/, 'Username must consist of 3-16 letters/numbers'),
            email: isEmail('Invalid email'),
            password: hasLength({min: 3}, 'Password must consist of at least 3 characters'),
        },
        validateInputOnBlur: true,
        clearInputErrorOnChange: true,
    });

    return (
        <Box
            component="form" className="mx-auto max-w-sm"
            onSubmit={form.onSubmit(values => signUp(values))}
        >
            <TextInput
                className="my-3"
                label="Email" placeholder="example@mail.com" withAsterisk
                {...form.getInputProps('email')}
            />
            <TextInput
                className="my-3"
                label="Username" placeholder="example" withAsterisk
                {...form.getInputProps('username')}
            />
            <PasswordInput
                className="my-3"
                label="Password" placeholder="****" withAsterisk
                {...form.getInputProps('password')}
            />

            <Group position="apart" className="mt-4 mb-2 -mx-1.5">
                <Button
                    size="sm" classNames={{label: 'text-xs font-medium'}} variant="subtle" compact
                    onClick={onSwitchToLogin}
                >I already have an account</Button>
            </Group>
            <Button
                type="submit" variant="filled" className="border-gray-800 hover:border-0 w-full"
                loading={isLoading}
            >Register</Button>
        </Box>
    );
};
