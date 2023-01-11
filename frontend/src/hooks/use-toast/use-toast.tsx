import React, { useCallback } from 'react';
import { showNotification } from '@mantine/notifications';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { faCheck, faExclamation, faInfo, faXmark } from '@fortawesome/free-solid-svg-icons';
import { MantineColor } from '@mantine/core';

type ToastCallback = (message: string, title?: string) => void;

export const useToast = (): {
    infoToast: ToastCallback,
    successToast: ToastCallback,
    errorToast: ToastCallback,
    warningToast: ToastCallback,
} => {

    const toast = useCallback((icon: IconProp, color: MantineColor): ToastCallback => {
        return (message: string, title?: string) => showNotification({
            title,
            message,
            icon: <FontAwesomeIcon icon={icon} />,
            color: color,
            classNames: {
                root: 'bg-white',
                description: 'text-dark-600',
                title: 'text-dark-800',
                closeButton: 'text-dark-600 hover:bg-gray-300',
            },
        });
    }, []);

    return {
        infoToast: toast(faInfo, 'violet'),
        successToast: toast(faCheck, 'green'),
        errorToast: toast(faXmark, 'red'),
        warningToast: toast(faExclamation, 'yellow'),
    };
};
