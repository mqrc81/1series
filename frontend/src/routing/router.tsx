import { lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';
import App from '../app';

const PopularShows = lazy(() => import('../pages/popular-shows/popular-shows'));
const UpcomingReleases = lazy(() => import('../pages/upcoming-releases/upcoming-releases'));

export const Router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
        children: [
            {
                path: '/shows/popular',
                element: <PopularShows/>,
            },
            {
                path: '/shows/releases',
                element: <UpcomingReleases/>,
            },
        ],
    },
]);
