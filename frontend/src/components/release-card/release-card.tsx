import React from 'react';
import { ReleaseDto } from '../../api';
import { Badge, Card, Group, Image, Text } from '@mantine/core';

type Props = {
    release: ReleaseDto;
}

export const ReleaseCard: React.FC<Props> = ({release}) => {
    return (
        <Card withBorder>
            <Card.Section>
                <Image
                    src={release.season.poster ?? release.show.poster}
                    withPlaceholder
                />
            </Card.Section>
            <Group position="apart" className="mt-5 mb-2">
                <Text className="font-semibold">{release.show.name + ' (' + release.season.name + ')'}</Text>
                <Badge color="pink" variant="light">
                    {release.show.rating}
                </Badge>
            </Group>
            <Text size="sm" color="dimmed">
                {release.season.overview ?? 'No overview for this season yet...'}
            </Text>
        </Card>
    );
};
