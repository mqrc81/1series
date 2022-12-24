import React, { Suspense } from 'react';
import AppLayout from './pages/layout/layout';
import { Spin } from 'antd';

const App: React.FC = () => {
    return (
        <div className="flex items-center">
            <Suspense fallback={<Spin className="m-auto mt-5" spinning size="large"/>}>
                <AppLayout/>
            </Suspense>
        </div>
    );
};

export default App;
