import React from 'react';
import { useGetShowQuery } from '../../api';
import { Loader } from '@mantine/core';
import { useParams } from 'react-router-dom';

const ShowDetails: React.FC = () => {
    const {id} = useParams();

    const {
        data: show,
        isSuccess, isLoading,
    } = useGetShowQuery(Number(id));

    return (
        <div className="w-full">
            {isLoading && (
                // TODO ms: skeleton
                <Loader className="m-auto mt-5"/>
            )}
            {isSuccess && (
                <div>{show.name}</div>
            )}
        </div>
    );
};

export default ShowDetails;
