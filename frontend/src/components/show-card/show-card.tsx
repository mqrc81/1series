import React from 'react';
import { ShowDto } from '../../api';
import { Badge, Card, Image, Text, Tooltip } from '@mantine/core';
import { Link } from 'react-router-dom';
import { useDebouncedState } from '@mantine/hooks';

type Props = {
    show: ShowDto;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    const [isHovering, setIsHovering] = useDebouncedState(false, 1200);
    return (
        <Link to={'/shows/' + show.id} className="cursor-pointer">
            <Card withBorder className="h-[414px]">
                <Card.Section className="relative" onMouseEnter={() => setIsHovering(true)}>
                    <Tooltip.Floating label={show.genres.map(({name}) => name).join(', ')} className={!isHovering && 'opacity-0'}>
                        <Image
                            src={show.poster}
                            withPlaceholder
                            className="relative"
                        />
                    </Tooltip.Floating>
                    <Badge color="violet" variant="gradient" className="absolute z-10 left-5 -bottom-3 py-3">
                        {show.rating.toFixed(1)}
                    </Badge>
                </Card.Section>
                <Card.Section className="mt-5 whitespace-nowrap">
                    <Text className="font-medium">{show.name}</Text>
                </Card.Section>
            </Card>
        </Link>
    );
};
