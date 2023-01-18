import React from 'react';
import { Button, Card } from '@mantine/core';
import { useGetGenresQuery } from '../../api';
import { useGenresFilter } from '../../hooks';
import { useForceUpdate } from '@mantine/hooks';

export const ShowFilters: React.FC<{ onFilterChange?: () => void }> = ({onFilterChange}) => {
    const forceUpdate = useForceUpdate();
    const {data: genres = []} = useGetGenresQuery();
    const {isGenreFiltered, filterGenre, unfilterGenre} = useGenresFilter();

    const onGenresChange = () => {
        forceUpdate();
        onFilterChange();
    };

    return (
        <Card withBorder>
            <div>
                Filter by genres
            </div>
            {genres
                .filter(isGenreFiltered)
                .map((genre, i) => (
                    <Button
                        radius="xl"
                        size="xs"
                        key={i}
                        color="teal"
                        className="cursor-pointer m-1 bg-teal-600 text-white"
                        variant="filled"
                        onClick={() => {
                            unfilterGenre(genre);
                            onGenresChange();
                        }}
                    >{genre.name}</Button>
                ))}
            {genres
                .filter(genre => !isGenreFiltered(genre))
                .map((genre, i) => (
                    <Button
                        radius="xl"
                        size="xs"
                        key={i}
                        className="cursor-pointer m-1 hover:bg-violet-600 hover:text-white"
                        variant="outline"
                        onClick={() => {
                            filterGenre(genre);
                            onGenresChange();
                        }}
                    >{genre.name}</Button>
                ))}
        </Card>
    );
};
