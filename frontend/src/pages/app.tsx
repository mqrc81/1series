import React, { Suspense } from 'react';
import Layout from './layout/layout';
import { usePingQuery } from '../api';
import { Loader } from '@mantine/core';

const App: React.FC = () => {
    usePingQuery();
    return (
        <Suspense fallback={<Loader color="teal" className="mx-auto mt-5 justify-self-center" />}>
            <Layout />;
        </Suspense>
    );
};

export default App;
