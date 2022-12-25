import { ThemeConfig } from 'antd/es/config-provider/context';
import { theme } from 'antd';

export const AntTheme: ThemeConfig = {
    algorithm: theme.darkAlgorithm,
    token: {
        colorTextBase: '#F3F4F6', // gray-100
        colorBgBase: '#374151', // gray-700
        colorPrimary: '#22D3EE', // cyan-400
        colorInfo: '#0EA5E9', // sky-500
        colorSuccess: '#22C55E', // green-500
        colorError: '#EF4444', // red-500
        colorWarning: '#EAB308', // yellow-500
        fontFamily: 'Rubik, sans-serif',
        // fontSize: 14, // default
        // borderRadius: 6, // default
        // lineType: 'solid', // default
        // lineWidth: 1, // default
        // controlHeight: 32, // default
        // zIndexBase: 0, // default
        // zIndexPopupBase: 1000, // default
        // sizeStep: 4, // default
        // sizeUnit: 4, // default
        // motionUnit: 0.1, // default
        // wireframe: false, // default
    },
};
