import React, { Suspense } from 'react';
import AppLayout from './layout/layout';
import { useQuery } from 'react-query';
import { GetQuery } from '../api';
import { Loader } from '@mantine/core';

const App: React.FC = () => {
    useQuery(['ping'], GetQuery('/ping'), {staleTime: Infinity});
    return (
        <Suspense fallback={<Loader className="mx-auto my-auto justify-self-center"/>}>
            <AppLayout/>
        </Suspense>
    );
};

export default App;
