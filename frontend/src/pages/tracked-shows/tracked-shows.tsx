import React from 'react';
import { useGetTrackedShowsQuery } from '../../api/users/users.queries';
import { Loader } from '@mantine/core';
import { ShowCard } from '../../components';
import ImportImdbWatchlist from '../import-imdb-watchlist/import-imdb-watchlist';

const TrackedShows: React.FC = () => {
    const {
        data: trackedShows = [],
        isSuccess, isLoading,
    } = useGetTrackedShowsQuery();

    return (<>
        <div className="grid grid-cols-1 gap-y-10">
            <div className="w-full">
                <div className="text-3xl">Tracked Series</div>
            </div>
            <div>
                {isLoading && (<Loader color="teal" className="m-auto mt-5" />)}
                {isSuccess && (
                    trackedShows.length > 0 ? (
                        <div className="grid grid-cols-4 md:grid-cols-7 gap-5">
                            {trackedShows.map((show, i) => (
                                <ShowCard key={i} show={show} />
                            ))}
                        </div>
                    ) : (
                        <div>You aren't tracking any shows yet!</div>
                    )
                )}
            </div>
            <div>
                <ImportImdbWatchlist />
            </div>
        </div>
    </>);
};

export default TrackedShows;
