import React from 'react';
import { Box, Button, Group, PasswordInput, TextInput } from '@mantine/core';
import { useToast } from '../../../hooks';
import { useAuthStore } from '../../../stores';
import { SignUserInDto, useSignUserInMutation } from '../../../api';
import { hasLength, isEmail, useForm } from '@mantine/form';

type Props = {
    onSubmit: () => void;
    onSwitchToRegister: () => void;
    onSwitchToResetPassword: () => void;
}

export const AuthModalLogin: React.FC<Props> = ({onSubmit, onSwitchToRegister, onSwitchToResetPassword}) => {
    const {successToast} = useToast();
    const {login} = useAuthStore();
    const {
        mutate: signIn,
        isLoading,
    } = useSignUserInMutation({
        onSuccess: (user) => {
            login(user);
            successToast('Successfully logged in!');
            onSubmit();
        },
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
        clearInputErrorOnChange: true,
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

            <Group position="apart" className="mt-4 mb-2 -mx-1.5">
                <Button
                    size="sm" classNames={{label: 'text-xs font-medium'}} variant="subtle" compact
                    onClick={onSwitchToRegister}
                >Create an account</Button>
                <Button
                    size="sm" classNames={{label: 'text-xs font-medium'}} variant="subtle" compact
                    onClick={onSwitchToResetPassword}
                >Forgot password</Button>
            </Group>
            <Button
                type="submit" variant="filled" className="border-gray-800 hover:border-0 w-full"
                loading={isLoading}
            >Log In</Button>
        </Box>
    );
};
