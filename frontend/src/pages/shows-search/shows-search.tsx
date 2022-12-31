import React from 'react';
import { useSearchShowsQuery } from '../../api';
import { BackgroundImage, Center, Loader } from '@mantine/core';
import { useSearchParams } from 'react-router-dom';

const ShowsSearch: React.FC = () => {
    const [searchParams] = useSearchParams('');

    const searchTerm = searchParams.get('searchTerm');

    const {
        data: searchResults,
        isSuccess, isLoading,
    } = useSearchShowsQuery(searchTerm, {minParamLength: 1});

    return (
        <div className="w-full">
            {isLoading && searchTerm.length > 0 && (
                // TODO ms: skeleton
                <Loader className="m-auto mt-5"/>
            )}
            {isSuccess && (
                <div className="grid grid-cols-3 gap-x-5 gap-y-3">
                    {searchResults.map((show, i) => (

                        <BackgroundImage src={show.poster} radius="md" className="h-[252px] w-[448]" key={i}>
                            <div key={i} className="bg-black/50 h-full flex">
                                <Center className="font-semibold text-3xl text-center mx-auto">{show.name}</Center>
                            </div>
                        </BackgroundImage>
                    ))}
                </div>
            )}
        </div>
    );
};

export default ShowsSearch;
