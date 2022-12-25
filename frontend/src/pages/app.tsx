import React, { Suspense } from 'react';
import AppLayout from './layout/layout';
import { Spin } from 'antd';

const App: React.FC = () => {
    return (
        <Suspense fallback={<Spin className="m-auto mt-5" spinning size="large"/>}>
            <AppLayout/>
        </Suspense>
    );
};

export default App;
