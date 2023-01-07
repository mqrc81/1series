import React from 'react';
import { Card, Skeleton } from '@mantine/core';

export const CardSkeleton: React.FC = () => {
    return (
        <Card className="h-[346px] w-[181px]">
            <Card.Section>
                <Skeleton height={272} />
            </Card.Section>
            <Skeleton height={24} className="mt-6" />
        </Card>
    );
};
