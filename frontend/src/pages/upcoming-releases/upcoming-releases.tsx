import React from 'react';
import { useGetUpcomingReleasesQuery } from '../../api';
import { useGenresFilter } from '../../hooks';
import { CardSkeleton, ReleaseCard, ShowFilters } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';
import { Button, Loader } from '@mantine/core';
import { useForceUpdate } from '@mantine/hooks';

const UpcomingReleases: React.FC = () => {
    const forceUpdate = useForceUpdate();
    const {
        data: releasesData,
        isSuccess, isLoading,
        fetchNextPage, hasNextPage,
        fetchPreviousPage, hasPreviousPage, isFetchingPreviousPage,
    } = useGetUpcomingReleasesQuery();
    const {isGenreFiltered, isGenreFilterActive} = useGenresFilter();

    return (
        <div className="grid grid-cols-1">
            <div className="flex flex-row">
                <div className="basis-1/4 mr-5 hidden md:block">
                    <ShowFilters onFilterChange={forceUpdate} />
                </div>
                <div>
                    {isSuccess && (
                        isFetchingPreviousPage || isLoading
                    ) ? (<>
                        <Loader color="teal" className="m-auto mt-5" />
                        <div className="grid grid-cols-3 md:grid-cols-5 gap-5 mt-5">
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <span className="hidden md:block"><CardSkeleton /></span>
                            <span className="hidden md:block"><CardSkeleton /></span>
                        </div>
                    </>) : (releasesData?.pages.length > 0 &&
                        <Button
                            className="mb-10 w-full"
                            disabled={!hasPreviousPage}
                            onClick={() => fetchPreviousPage()}
                        >Load past shows</Button>
                    )}
                    {isSuccess && (
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
                            dataLength={releasesData?.pages.length}
                        >
                            <div className="grid grid-cols-5 gap-5">
                                {releasesData?.pages
                                    .flatMap(({releases}) => releases)
                                    .filter(({show: {genres}}) => !isGenreFilterActive() || genres.some(isGenreFiltered))
                                    .map((release, i) => (
                                        <ReleaseCard key={i} release={release} />
                                    ))}
                            </div>
                        </InfiniteScroll>
                    )}
                </div>
            </div>
        </div>
    );
};

export default UpcomingReleases;
