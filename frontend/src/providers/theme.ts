import { MantineThemeOverride } from '@mantine/core';

export const MantineTheme: MantineThemeOverride = {
    colorScheme: 'dark',
    fontFamily: 'Rubik, sans-serif',
    primaryColor: 'violet',
    primaryShade: 6,
    components: {
        Loader: {
            defaultProps: {
                size: 'xl',
            },
        },
        Card: {
            defaultProps: {
                shadow: 'md',
                radius: 'md',
            },
        },
        Image: {
            defaultProps: {
                radius: 'md',
            },
        },
    },
    defaultRadius: 'md',
    cursorType: 'pointer',
    loader: 'oval',
    defaultGradient: {from: 'violet', to: 'teal'},
};
