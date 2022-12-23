import { message } from 'antd';
import React, { useCallback } from 'react';
import { NoticeType } from 'antd/es/message/interface';

type ToastCallback = (text: string) => void;

export const useToast = (): {
    infoToast: ToastCallback,
    successToast: ToastCallback,
    errorToast: ToastCallback,
    warningToast: ToastCallback,
    toastContextHolder: ReturnType<typeof message['useMessage']>[1]
} => {
    const [messageApi, contextHolder] = message.useMessage({
        maxCount: 3,
    });

    const toast = useCallback((type: NoticeType): ToastCallback =>
        (text: string) => messageApi.open({
            type,
            content: text,
        }), []);

    return {
        infoToast: toast('info'),
        successToast: toast('success'),
        errorToast: toast('error'),
        warningToast: toast('warning'),
        toastContextHolder: contextHolder,
    };
};
