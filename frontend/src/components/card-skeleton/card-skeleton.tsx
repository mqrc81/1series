import React from 'react';
import { Card, Skeleton } from '@mantine/core';

export const CardSkeleton: React.FC = () => {
    return (
        <Card className="h-[414px]">
            <Card.Section>
                <Skeleton height={352}/>
            </Card.Section>
            <Card.Section className="px-2 py-5">
                <Skeleton height={20}/>
            </Card.Section>
        </Card>
    );
};
