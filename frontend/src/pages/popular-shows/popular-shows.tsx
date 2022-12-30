import React from 'react';
import { useGetPopularShowsQuery } from '../../api';
import { useToast } from '../../hooks/use-toast/use-toast';
import { ShowCard } from '../../components';
import InfiniteScroll from 'react-infinite-scroll-component';
import { Loader } from '@mantine/core';

const PopularShows: React.FC = () => {
    const {errorToast} = useToast();
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
                >
                    <div className="grid grid-cols-5 gap-5">
                        {showsData.pages.flatMap(({shows}) => shows).map((show, i) => (
                            <ShowCard key={i} show={show}/>
                        ))}
                    </div>
                </InfiniteScroll>
            }
            {(isLoading || isFetching) && <Loader className="m-auto mt-5"/>}
        </div>
    );
};

export default PopularShows;
