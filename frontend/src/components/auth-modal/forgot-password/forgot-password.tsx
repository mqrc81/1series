import React from 'react';
import { useToast } from '../../../hooks';
import { ForgotPasswordDto, useForgotPasswordMutation } from '../../../api';
import { isEmail, useForm } from '@mantine/form';
import { Box, Button, Group, TextInput } from '@mantine/core';

type Props = {
    onSubmit: () => void;
    onSwitchToLogin: () => void;
}

export const AuthModalForgotPassword: React.FC<Props> = ({onSubmit, onSwitchToLogin}) => {
    const {successToast, errorToast} = useToast();
    const {
        mutate: requestPasswordReset,
        isLoading,
    } = useForgotPasswordMutation({
        onSuccess: () => {
            successToast('A link to reset your password has been sent to the email!');
            onSubmit();
        },
        onError: () => errorToast('A user with this email doesn\'t exist!'),
    });

    const form = useForm<ForgotPasswordDto>({
        initialValues: {
            email: '',
        },
        validate: {
            email: isEmail('Invalid email'),
        },
        validateInputOnBlur: true,
        clearInputErrorOnChange: true,
    });

    return (
        <Box
            component="form" className="mx-auto max-w-sm"
            onSubmit={form.onSubmit(values => requestPasswordReset(values))}
        >
            <TextInput
                className="my-3"
                label="Email" placeholder="example@mail.com" withAsterisk
                {...form.getInputProps('email')}
            />

            <Group position="apart" className="mt-4 mb-2 -mx-1.5">
                <Button
                    size="sm" classNames={{label: 'text-xs font-medium'}} variant="subtle" compact
                    onClick={onSwitchToLogin}
                >Back to Login</Button>
            </Group>
            <Button
                type="submit" variant="filled" className="border-gray-800 hover:border-0 w-full"
                loading={isLoading}
            >Request Password Reset</Button>
        </Box>
    );
};
