import React, { useCallback } from 'react';
import { showNotification } from '@mantine/notifications';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { faCheck, faExclamation, faInfo, faXmark } from '@fortawesome/free-solid-svg-icons';

type ToastCallback = (message: string, title?: string) => void;

export const useToast = (): {
    infoToast: ToastCallback,
    successToast: ToastCallback,
    errorToast: ToastCallback,
    warningToast: ToastCallback,
} => {

    const toast = useCallback((icon: IconProp, color: string): ToastCallback => {
        return (message: string, title?: string) => showNotification({
            title,
            message,
            icon: <FontAwesomeIcon icon={icon} className={'text-' + color + '-600'} />,
        });
    }, []);

    return {
        infoToast: toast(faInfo, 'violet'),
        successToast: toast(faCheck, 'green'),
        errorToast: toast(faXmark, 'red'),
        warningToast: toast(faExclamation, 'yellow'),
    };
};
