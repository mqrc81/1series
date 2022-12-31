import React from 'react';
import { ShowDto } from '../../api';
import { Badge, Card, Image, Text } from '@mantine/core';
import { Link } from 'react-router-dom';

type Props = {
    show: ShowDto;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    return (
        <Link to={'/shows/' + show.id} className="cursor-pointer">
            <Card withBorder className="h-[414px]">
                <Card.Section className="relative">
                    <Image
                        src={show.poster}
                        withPlaceholder
                        className="relative"
                    />
                    <Badge color="violet" variant="gradient" className="absolute z-10 left-5 -bottom-3 py-3">
                        {show.rating}
                    </Badge>
                </Card.Section>
                <Card.Section className="mt-5 whitespace-nowrap">
                    <Text className="font-semibold">{show.name}</Text>
                </Card.Section>
            </Card>
        </Link>
    );
};
