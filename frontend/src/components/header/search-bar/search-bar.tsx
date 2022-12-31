import React, { forwardRef, useRef, useState } from 'react';
import { Divider, Group, Image, Loader, Select } from '@mantine/core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
import { SelectItem } from '@mantine/core/lib/Select/types';
import { useDebouncedValue } from '@mantine/hooks';
import { ShowSearchResultDto, useSearchShowsQuery } from '../../../api';

type SearchShowsData = SelectItem & { group: 'Series' | 'Other', onSelect: () => void } & ShowSearchResultDto;

const SearchResult = forwardRef<HTMLDivElement, SearchShowsData>((data, ref) => {
        const {id, name, poster, ...other} = data;
        return (
            <>
                <div ref={ref} {...other} className="p-2 aria-selected:bg-gray-700 cursor-pointer">
                    <Group noWrap>
                        <Image withPlaceholder width={50} height={75} src={poster} radius="md"/>
                        <div>
                            <div className="font-medium mb-auto">
                                {name}
                            </div>
                        </div>
                    </Group>
                </div>
                <Divider/>
            </>
        );
    },
);

export const HeaderSearchBar: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [debouncedSearchTerm] = useDebouncedValue(searchTerm, 300);
    const navigate = useNavigate();

    const {
        data: searchResults = [],
        isLoading,
    } = useSearchShowsQuery(debouncedSearchTerm, {minParamLength: 3});

    const selectRef = useRef<HTMLInputElement>();

    const navigateToSearchPage = () => {
        if (searchTerm.length < 3) return;
        navigate('/shows/search?searchTerm=' + searchTerm);
        resetSearchInput();
    };
    const navigateToShowPage = (showId: string) => {
        if (searchTerm.length < 3) return;
        navigate('/shows/' + showId);
        resetSearchInput();
    };
    const resetSearchInput = () => {
        selectRef.current.blur();
        setSearchTerm('');
    };

    return (
        <Select
            ref={selectRef}
            className="w-[28rem]"
            placeholder="Search series"
            rightSection={isLoading ? (
                <Loader size="xs"/>
            ) : (
                <FontAwesomeIcon
                    icon={faMagnifyingGlass}
                    onClick={navigateToSearchPage}
                    className="cursor-pointer text-violet-600"
                />
            )}
            itemComponent={SearchResult}
            data={searchResults.map((show) => ({
                value: show.id + '',
                group: 'Series',
                label: show.name,
                ...show,
            } as SearchShowsData))}
            searchable
            onChange={(showId) => navigateToShowPage(showId)}
            searchValue={searchTerm}
            onInput={({currentTarget: {value}}) => setSearchTerm(value)}
            nothingFound={(searchTerm === '' || isLoading) ? null : (searchTerm.length < 3 ? 'Type at least 3 characters' : 'No series found')}
            maxDropdownHeight={300}
            classNames={{input: 'bg-white text-violet-600'}}
            onKeyDown={({key}) => {if (key === 'Enter') navigateToSearchPage();}}
        />
    );
};
