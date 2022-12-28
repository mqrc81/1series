import React, { Suspense } from 'react';
import AppLayout from './layout/layout';
import { Spin } from 'antd';
import { useQuery } from 'react-query';
import { GetQuery } from '../api';

const App: React.FC = () => {
    useQuery(['ping'], GetQuery('/ping'), {cacheTime: Infinity});
    return (
        <Suspense fallback={<Spin className="m-auto mt-5" spinning size="large"/>}>
            <AppLayout/>
        </Suspense>
    );
};

export default App;
