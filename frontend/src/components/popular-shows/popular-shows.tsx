import React, { useState } from 'react';
import { useGetPopularShowsQuery } from '../../api';
import { Card, Image, Spin } from 'antd';

const PopularShows: React.FC = () => {
    const [page, setPage] = useState(1);
    const {data: shows, isLoading, isSuccess} = useGetPopularShowsQuery(page);
    return (
        <>
            <Spin spinning={isLoading} size="large"/>
            {isSuccess && <div className="grid grid-cols-4 gap-4">
                {shows?.map(show => (
                    <Card bordered title={show.name}>
                        <div className="text-primary block">
                            <div className="mb-4">{show.overview}</div>
                            <div className="w-3/4 mx-auto">
                                <Image src={show.poster}/>
                            </div>
                        </div>
                        <br/>
                    </Card>
                ))}
            </div>}
        </>
    );
};

export default PopularShows;
