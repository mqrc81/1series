import React from 'react';
import { ReleaseDto } from '../../api';
import { Badge, Card, Image, Text } from '@mantine/core';
import { Link } from 'react-router-dom';

type Props = {
    release: ReleaseDto;
}

export const ReleaseCard: React.FC<Props> = ({release: {show, season}}) => {
    return (
        <Link to={'/shows/' + show.id} className="cursor-pointer">
            <Card withBorder className="h-[414px]">
                <Card.Section className="relative">
                    <Image
                        src={season.poster ?? show.poster}
                        withPlaceholder
                        height={352}
                    />
                    {!!show.rating && <Badge color="violet" variant="gradient" className="absolute z-10 left-5 -bottom-3 py-3">
                        {show.rating}
                    </Badge>}
                    <div className="bg-black/50 text-white absolute z-10 top-0 right-0 text-sm rounded-md px-2 py-1">{season.name}</div>
                </Card.Section>
                <Card.Section className="h-62 px-2 py-5">
                    <Text className="font-semibold">{show.name}</Text>
                </Card.Section>
            </Card>
        </Link>
    );
};
