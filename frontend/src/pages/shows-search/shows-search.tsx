import React from 'react';
import { useSearchShowsQuery } from '../../api';
import { BackgroundImage, Card, Center, Loader } from '@mantine/core';
import { Link, useSearchParams } from 'react-router-dom';

const ShowsSearch: React.FC = () => {
    const [searchParams] = useSearchParams('');
    const searchTerm = searchParams.get('searchTerm');

    const {
        data: searchResults,
        isSuccess, isLoading,
    } = useSearchShowsQuery(searchTerm, {minParamLength: 3});

    return (
        <div className="w-full">
            {isLoading && (
                // TODO ms: skeleton?
                <Loader className="m-auto mt-5"/>
            )}
            {isSuccess && (
                <div className="grid grid-cols-3 gap-x-5 gap-y-3">
                    {searchResults.map((show, i) => (
                        <Link to={'/shows/' + show.id} key={i} className="cursor-pointer">
                            <Card radius="md">
                                <Card.Section>
                                    <BackgroundImage src={show.poster} className="h-[252px] w-[448]">
                                        <div key={i} className="bg-black/50 h-full flex">
                                            <Center className="font-semibold text-3xl text-center mx-auto">{show.name}</Center>
                                        </div>
                                    </BackgroundImage>
                                </Card.Section>
                            </Card>
                        </Link>
                    ))}
                </div>
            )}
        </div>
    );
};

export default ShowsSearch;
