import React from 'react';
import { useGetPopularShowsQuery } from '../../api';
import { Spin } from 'antd';
import { useToast } from '../../hooks/use-toast/use-toast';
import { ShowCard } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';

const PopularShows: React.FC = () => {
    const {errorToast, toastContextHolder} = useToast();
    const {
        data: showsData,
        isSuccess, isLoading, isFetching,
        fetchNextPage, hasNextPage,
    } = useGetPopularShowsQuery({
        onError: () => errorToast('Error fetching popular shows...'),
    });

    return (
        <div className="grid grid-cols-1 w-full">
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
                        {showsData.pages.flatMap(({shows}) => shows).map((show, i) => (
                            <ShowCard key={i} show={show}/>
                        ))}
                    </div>
                </InfiniteScroll>
            }
            {(isLoading || isFetching) && <Spin className="m-auto mt-5" spinning size="large"/>}
            {toastContextHolder}
        </div>
    );
};

export default PopularShows;
