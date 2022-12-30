import React from 'react';
import { useGetUpcomingReleasesQuery } from '../../api';
import { useToast } from '../../hooks/use-toast/use-toast';
import { ReleaseCard } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';
import { Button, Loader } from '@mantine/core';

const UpcomingReleases: React.FC = () => {
    const {errorToast} = useToast();
    const {
        data: showsData,
        isSuccess, isLoading,
        fetchNextPage, hasNextPage, isFetchingNextPage,
        fetchPreviousPage, hasPreviousPage, isFetchingPreviousPage,
    } = useGetUpcomingReleasesQuery({
        onError: () => errorToast('Error fetching upcoming releases...'),
    });

    return (
        <div className="grid grid-cols-1 w-full">
            {isSuccess && (
                isFetchingPreviousPage || isLoading
            ) ? (
                <Loader className="m-auto mt-5"/>
            ) : (
                <Button
                    className="mb-10"
                    disabled={!hasPreviousPage}
                    onClick={() => fetchPreviousPage()}
                >Load past shows</Button>
            )}
            {isSuccess &&
                <InfiniteScroll
                    style={{overflow: 'hidden'}}
                    next={fetchNextPage}
                    hasMore={hasNextPage}
                    loader={undefined}
                    dataLength={showsData.pages.length}
                    scrollThreshold={0.9}
                >
                    <div className="grid grid-cols-5 gap-5">
                        {showsData.pages.flatMap(({releases}) => releases).map((release, i) => (
                            <ReleaseCard key={i} release={release}/>
                        ))}
                    </div>
                </InfiniteScroll>
            }
            {(isLoading || isFetchingNextPage) && <Loader className="m-auto mt-5"/>}
        </div>
    );
};

export default UpcomingReleases;
