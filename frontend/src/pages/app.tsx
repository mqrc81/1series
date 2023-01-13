import React, { Suspense } from 'react';
import Layout from './layout/layout';
import { useInitQuery, User } from '../api';
import { Loader } from '@mantine/core';
import { useAuthStore } from '../stores';

const App: React.FC = () => {
    const {login} = useAuthStore();
    useInitQuery({
        onSuccess: (user?: User) => {
            if (user) {
                login(user);
            }
        },
    });
    return (
        <Suspense fallback={<Loader color="teal" className="mx-auto mt-5 justify-self-center" />}>
            <Layout />;
        </Suspense>
    );
};

export default App;
