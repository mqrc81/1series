import React from 'react';
import { Link, useRouteError } from 'react-router-dom';
import { Button } from '@mantine/core';

type Props = { title: string, subtitle: string, description: string };

const ErrorPage: React.FC<Props> = ({title, subtitle, description}) => {
    return (
        <div className="h-screen w-screen flex bg-gray-700">
            <div className="m-auto max-w-screen-xl text-center">
                <h1 className="mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-violet-600">{title}</h1>
                <p className="mb-4 text-3xl tracking-tight font-bold md:text-4xl text-white">{subtitle}</p>
                <p className="mb-4 text-lg font-light text-gray-400">{description}</p>
                <Button className="inline-flex" size="lg">
                    <Link className="hover:text-violet-600" to="/">Back to Homepage</Link>
                </Button>
            </div>
        </div>
    );
};

const ErrorFallback: React.FC = () => {
    const error = useRouteError() as { status?: number };

    return (
        error?.status === 404
    ) ? (
        <ErrorPage
            title="404"
            subtitle="Something's missing."
            description="Sorry, we can't find that page. You'll find lots to explore on the home-page."
        />
    ) : (
        <ErrorPage
            title="Error"
            subtitle="Something's wrong."
            description="Sorry, an unknown error occurred. Please try again by visiting the home-page."
        />
    );
};

export default ErrorFallback;
