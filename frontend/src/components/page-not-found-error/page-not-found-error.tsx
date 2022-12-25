import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'antd';

export const PageNotFoundError: React.FC = () => {
    return (
        <div className="h-screen flex bg-gray-700">
            <div className="m-auto max-w-screen-xl text-center">
                <h1 className="mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-cyan-400">404</h1>
                <p className="mb-4 text-3xl tracking-tight font-bold md:text-4xl text-white">Something's missing.</p>
                <p className="mb-4 text-lg font-light text-gray-400">Sorry, we can't find that page. You'll find lots to explore on the home page.</p>
                <Button className="inline-flex" size="large">
                    <Link className="hover:text-cyan-400" to="/">Back to Homepage</Link>
                </Button>
            </div>
        </div>
    );
};
