import React, { lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';
import App from '../pages/app';
import ErrorFallback from '../pages/error-fallback/error-fallback';
import { ApiClient } from './apisauce';

const Home = lazy(() => import('../pages/home/home'));
const PopularShows = lazy(() => import('../pages/popular-shows/popular-shows'));
const UpcomingReleases = lazy(() => import('../pages/upcoming-releases/upcoming-releases'));
const ImportImdbWatchlist = lazy(() => import('../pages/import-imdb-watchlist/import-imdb-watchlist'));

export const Router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
        errorElement: <ErrorFallback/>,
        loader: () => ApiClient.get('/ping', {}, {withCredentials: false}),
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
                path: '/users/import-imdb-watchlist',
                element: <ImportImdbWatchlist/>,
            },
        ],
    },
]);
