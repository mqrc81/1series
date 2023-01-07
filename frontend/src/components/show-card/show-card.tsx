import React from 'react';
import { Show } from '../../api';
import { Badge, Card, Image, Tooltip } from '@mantine/core';
import { Link } from 'react-router-dom';

type Props = {
    show: Show;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    return (
        <Link to={'/shows/' + show.id} className="cursor-pointer">
            <Card withBorder shadow="xl">
                <Card.Section className="relative">
                    <Image src={show.poster} withPlaceholder className="h-[272px]" />
                    {!!show.rating && (
                        <Badge variant="gradient" className="absolute z-10 left-4 -bottom-3 py-3">
                            {show.rating.toFixed(1)}
                        </Badge>
                    )}
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
        </Link>
    );
};
