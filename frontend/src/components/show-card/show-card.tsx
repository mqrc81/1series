import React from 'react';
import { ShowDto } from '../../api';
import { Badge, Card, Group, Image, Text } from '@mantine/core';

type Props = {
    show: ShowDto;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    return (
        <>
            <Card withBorder>
                <Card.Section>
                    <Image
                        src={show.poster}
                        withPlaceholder
                    />
                </Card.Section>
                <Group position="apart" className="mt-5 mb-2">
                    <Text className="font-semibold">{show.name}</Text>
                    <Badge color="pink" variant="light">
                        {show.rating}
                    </Badge>
                </Group>
                <Text size="sm" color="dimmed">
                    {show.overview ?? 'No overview for this series yet...'}
                </Text>
            </Card>
        </>
    );
};
