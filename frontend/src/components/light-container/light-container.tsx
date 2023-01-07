import React from 'react';
import { MantineProvider } from '@mantine/core';

export const LightContainer: React.FC<{ children: React.ReactNode }> = ({children}) => {
    return (
        <MantineProvider inherit theme={{colorScheme: 'light'}}>
            {children}
        </MantineProvider>
    );
};
