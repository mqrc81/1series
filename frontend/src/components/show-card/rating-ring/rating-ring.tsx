import React from 'react';
import { MantineColor, RingProgress, Tooltip } from '@mantine/core';

type Props = {
    rating: number | undefined;
    ratingCount: number;
    minRatingCount?: number;
    className?: string;
}

const color = (value: number): MantineColor => {
    if (value < 3) return 'red';
    if (value < 7) return 'yellow';
    if (value < 9) return 'teal';
    return 'violet';
};

const ROUNDED_CAPS_SPACE = 0.2;

export const ShowCardRatingRing: React.FC<Props> = ({rating, ratingCount, minRatingCount = 50, className = ''}) => {
    if (ratingCount < minRatingCount) {
        return <></>;
    }
    return (
        <div className={className}>
            <Tooltip
                label={ratingCount + ' ratings'}
                classNames={{tooltip: 'text-xs opacity-90'}}
                openDelay={500}
                position="right-start"
                offset={1}
                withArrow
                arrowPosition="center"
            >
                <RingProgress
                    sections={[{value: rating * (10 - ROUNDED_CAPS_SPACE), color: color(rating)}]}
                    roundCaps={true}
                    className="bg-dark-600/75 rounded-full"
                    label={
                        <div className="font-semibold text-center text-xs">
                            {rating.toFixed(1)}
                        </div>
                    }
                    thickness={3}
                    size={45}
                />
            </Tooltip>
        </div>
    );
};
