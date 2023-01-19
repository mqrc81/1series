import React from 'react';
import { Show } from '../../api';
import { Card, Image, Tooltip } from '@mantine/core';
import { useNavigate } from 'react-router-dom';
import { ShowCardActionsDropdown } from './actions-dropdown/actions-dropdown';
import { ShowCardRatingRing } from './rating-ring/rating-ring';

type Props = {
    show: Show;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    const navigate = useNavigate();
    return (
        <Card withBorder shadow="xl" onClick={() => navigate('/shows/' + show.id)} className="cursor-pointer">
            <Card.Section className="relative">
                <ShowCardActionsDropdown showId={show.id} className="absolute z-10 right-0 top-0 w-4" />
                <Image src={show.poster} withPlaceholder className="h-[272px]" />
                <ShowCardRatingRing
                    rating={show.rating} ratingCount={show.ratingCount}
                    className="absolute z-10 left-1 -bottom-5"
                />
            </Card.Section>
            <Tooltip
                label={show.genres.map(({name}) => name).join(', ')}
                openDelay={800}
                multiline position="top-end"
                classNames={{tooltip: 'opacity-80 text-sm'}}
            >
                <Card.Section className="p-2 pt-5 whitespace-nowrap h-14">
                    <span className="font-medium">{show.name}</span>
                </Card.Section>
            </Tooltip>
        </Card>
    );
};
