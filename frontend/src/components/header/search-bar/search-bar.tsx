import React, { forwardRef, useRef, useState } from 'react';
import { Divider, Group, Image, Loader, Select } from '@mantine/core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { NavLink } from 'react-router-dom';
import { SelectItem } from '@mantine/core/lib/Select/types';
import { useDebouncedValue } from '@mantine/hooks';
import { ShowSearchResultDto, useSearchShowsQuery } from '../../../api';

type SearchShowsData = SelectItem & { group: 'Series' | 'Other', onSelect: () => void } & ShowSearchResultDto;

const SearchResult = forwardRef<HTMLDivElement, SearchShowsData>((data, ref) => {
        const {id, name, poster, onSelect} = data;
        return (
            <>
                <div ref={ref} className="p-2 hover:bg-gray-700">
                    <NavLink to={'/shows/' + id} onClick={onSelect}>
                        <Group noWrap>
                            <Image withPlaceholder width={50} height={75} src={poster} radius="md"/>
                            <div>
                                <div className="font-semibold mb-auto">
                                    {name}
                                </div>
                            </div>
                        </Group>
                    </NavLink>
                </div>
                <Divider/>
            </>
        );
    },
);

export const HeaderSearchBar: React.FC = () => {
    const [searchInput, setSearchInput] = useState('');
    const [debounced] = useDebouncedValue(searchInput, 300);

    const {
        data: searchResults = [],
        isLoading,
    } = useSearchShowsQuery(debounced, {minParamLength: 3});

    const selectRef = useRef<HTMLInputElement>();

    return (
        <Select
            ref={selectRef}
            className="w-[28rem]"
            placeholder="Search series..."
            rightSection={isLoading ? <Loader size="xs"/> : <FontAwesomeIcon icon={faMagnifyingGlass}/>}
            itemComponent={SearchResult}
            data={searchResults.map((show) => ({
                value: show.id + '',
                group: 'Series',
                label: show.name,
                onSelect: () => {
                    selectRef.current.blur();
                    setSearchInput('');
                },
                ...show,
            } as SearchShowsData))}
            searchable
            searchValue={searchInput}
            onInput={(input) => setSearchInput(input.currentTarget.value)}
            nothingFound={(searchInput === '' || isLoading) ? null : (searchInput.length < 3 ? 'Type at least 3 characters' : 'No series found')}
            maxDropdownHeight={300}
            classNames={{
                input: 'bg-white text-violet-500',
            }}
        />
    );
};
