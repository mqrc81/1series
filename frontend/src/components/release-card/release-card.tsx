import React from 'react';
import { ReleaseDto } from '../../api';
import { Badge, Card, Image, Tooltip } from '@mantine/core';
import { Link } from 'react-router-dom';

type Props = {
    release: ReleaseDto;
}

export const ReleaseCard: React.FC<Props> = ({release: {show, season}}) => {
    return (
        <Link to={'/shows/' + show.id} className="cursor-pointer">
            <Card withBorder shadow="xl">
                <Card.Section className="relative">
                    <Image src={season.poster ?? show.poster} withPlaceholder className="h-[272px]" />
                    {!!show.rating && (
                        <Badge variant="gradient" className="absolute z-10 left-4 -bottom-3 py-3">
                            {show.rating.toFixed(1)}
                        </Badge>
                    )}
                    <div
                        className="bg-black/50 text-white absolute z-10 top-0 right-0 text-sm rounded-md px-2 py-1"
                    >{season.name}</div>
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
