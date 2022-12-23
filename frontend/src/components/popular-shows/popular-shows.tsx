import React, { useState } from 'react';
import { useGetPopularShowsQuery } from '../../api';
import { Spin } from 'antd';
import { useToast } from '../../hooks/use-toast/use-toast';
import { ShowCard } from '../elements';

const PopularShows: React.FC = () => {
    const {errorToast, toastContextHolder} = useToast();
    const [page, setPage] = useState(1);
    const {data: shows = [], isLoading, isSuccess} = useGetPopularShowsQuery(page, {
        onError: () => errorToast('Error fetching popular shows...'),
    });

    return (
        <>
            {toastContextHolder}
            <Spin className="m-auto" spinning={isLoading} size="large"/>
            {isSuccess && <div className="grid grid-cols-4 gap-5 m-10">
                {shows.map((show, i) => (
                    <ShowCard key={i} show={show}/>
                ))}
            </div>}
        </>
    );
};

export default PopularShows;
