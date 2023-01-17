import React from 'react';
import { useGetPopularShowsQuery } from '../../api';
import { useGenresFilter } from '../../hooks';
import { CardSkeleton, ShowCard, ShowFilters } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';
import { Loader } from '@mantine/core';
import { useForceUpdate } from '@mantine/hooks';

const PopularShows: React.FC = () => {
    const forceUpdate = useForceUpdate();
    const {
        data: showsData,
        fetchNextPage, hasNextPage,
    } = useGetPopularShowsQuery();
    const {isGenreFiltered, isGenreFilterActive} = useGenresFilter();

    return (
        <div className="grid grid-cols-1">
            <div className="flex flex-row">
                <div className="basis-1/4 mr-5 hidden md:block">
                    <ShowFilters onFilterChange={forceUpdate} />
                </div>
                <InfiniteScroll
                    style={{overflow: 'hidden'}}
                    next={fetchNextPage}
                    hasMore={hasNextPage}
                    loader={(<>
                        <div className="grid grid-cols-3 md:grid-cols-5 gap-5 mt-5">
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <span className="hidden md:block"><CardSkeleton /></span>
                            <span className="hidden md:block"><CardSkeleton /></span>
                        </div>
                        <Loader color="teal" className="m-auto mt-5" />
                    </>)}
                    dataLength={showsData?.pages.length ?? 0}
                >
                    <div className="grid grid-cols-3 md:grid-cols-5 gap-5">
                        {showsData?.pages
                            .flatMap(({shows}) => shows)
                            .filter(({genres}) => !isGenreFilterActive() || genres.some(isGenreFiltered))
                            .map((show, i) => (
                                <ShowCard key={i} show={show} />
                            ))}
                    </div>
                </InfiniteScroll>
            </div>
        </div>
    );
};

export default PopularShows;
