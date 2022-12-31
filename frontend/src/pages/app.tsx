import React, { Suspense } from 'react';
import Layout from './layout/layout';
import { usePingQuery } from '../api';
import { Loader } from '@mantine/core';

const App: React.FC = () => {
    usePingQuery();
    return (
        <Suspense fallback={<Loader className="mx-auto my-auto justify-self-center"/>}>
            <Layout/>
        </Suspense>
    );
};

export default App;
