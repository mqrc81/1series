import React from 'react';
import { useRouteError } from 'react-router-dom';
import { PageNotFoundError } from '../../components/page-not-found-error/page-not-found-error';
import { UnknownError } from '../../components/unknown-error/unknown-error';

const ErrorFallback: React.FC = () => {
    const error = useRouteError() as { status?: number };
    return (
        (error?.status === 404) ? (<PageNotFoundError/>) : (<UnknownError/>)
    );
};

export default ErrorFallback;
