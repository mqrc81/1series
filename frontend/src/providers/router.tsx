import React, { lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';
import App from '../pages/app';
import ErrorFallback from '../pages/error-fallback/error-fallback';

const Home = lazy(() => import('../pages/home/home'));
const PopularShows = lazy(() => import('../pages/popular-shows/popular-shows'));
const UpcomingReleases = lazy(() => import('../pages/upcoming-releases/upcoming-releases'));
const ShowDetails = lazy(() => import('../pages/show-details/show-details'));
const ShowsSearch = lazy(() => import('../pages/shows-search/shows-search'));

export const Router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
        errorElement: <ErrorFallback/>,
        children: [
            {
                path: '/',
                element: <Home/>,
            },
            {
                path: '/shows/popular',
                element: <PopularShows/>,
            },
            {
                path: '/shows/releases',
                element: <UpcomingReleases/>,
            },
            {
                path: '/shows/:id',
                element: <ShowDetails/>,
            },
            {
                path: '/shows/search',
                element: <ShowsSearch/>,
            },
        ],
    },
]);
