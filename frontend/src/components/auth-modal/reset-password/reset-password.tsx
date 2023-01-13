import React from 'react';
import { useToast } from '../../../hooks';
import { ResetPasswordDto, useResetPasswordMutation } from '../../../api';
import { hasLength, useForm } from '@mantine/form';
import { Box, Button, PasswordInput } from '@mantine/core';
import { useSearchParams } from 'react-router-dom';

type Props = {
    onSubmit: () => void;
}

export const AuthModalResetPassword: React.FC<Props> = ({onSubmit}) => {
    const {successToast, errorToast} = useToast();
    const {
        mutate: resetPassword,
        isLoading,
    } = useResetPasswordMutation({
        onSuccess: () => {
            successToast('Your password has been reset! Please log in.');
            onSubmit();
        },
        onError: () => errorToast('Your link has expired!'),
        onSettled: () => searchParams.delete('token'),
    });

    const [searchParams] = useSearchParams({token: ''});
    const token = searchParams.get('token');

    const form = useForm<ResetPasswordDto>({
        initialValues: {
            password: '',
        },
        validate: {
            password: hasLength({min: 3}, 'Password must consist of at least 3 characters'),
        },
        validateInputOnBlur: true,
        clearInputErrorOnChange: true,
    });

    return (
        <Box
            component="form" className="mx-auto max-w-sm"
            onSubmit={form.onSubmit(values => resetPassword({...values, token}))}
        >
            <PasswordInput
                className="my-3"
                label="New Password" placeholder="****" withAsterisk
                {...form.getInputProps('password')}
            />

            <Button
                type="submit" variant="filled" className="border-gray-800 hover:border-0 w-full mt-4"
                loading={isLoading}
            >Reset Password</Button>
        </Box>
    );
};
