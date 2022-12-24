import React from 'react';
import { useGetUpcomingReleasesQuery } from '../../api';
import { Button, Spin } from 'antd';
import { useToast } from '../../hooks/use-toast/use-toast';
import { ReleaseCard } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';

const UpcomingReleases: React.FC = () => {
    const {errorToast, toastContextHolder} = useToast();
    const {
        data: showsData,
        isSuccess, isLoading, isFetching,
        fetchNextPage, hasNextPage,
        fetchPreviousPage, hasPreviousPage,
    } = useGetUpcomingReleasesQuery({
        onError: () => errorToast('Error fetching upcoming releases...'),
    });

    return (
        <div className="grid grid-cols-1 w-full">
            {isSuccess && <Button
                disabled={!hasPreviousPage || isFetching || isLoading}
                onClick={() => fetchPreviousPage()}
            >Load past shows</Button>}
            {isSuccess &&
                <InfiniteScroll
                    style={{overflow: 'hidden'}}
                    next={fetchNextPage}
                    hasMore={hasNextPage}
                    loader={undefined}
                    dataLength={showsData.pages.length}
                    scrollThreshold={0.9}
                >
                    <div className="grid grid-cols-4 gap-5 m-10">
                        {showsData.pages.flatMap(({releases}) => releases).map((release, i) => (
                            <ReleaseCard key={i} release={release}/>
                        ))}
                    </div>
                </InfiniteScroll>
            }
            {(isLoading || isFetching) && <Spin className="m-auto mt-5" spinning size="large"/>}
            {toastContextHolder}
        </div>
    );
};

export default UpcomingReleases;
