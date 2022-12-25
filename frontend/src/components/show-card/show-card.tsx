import React from 'react';
import { Card, Image } from 'antd';
import { ShowDto } from '../../api';

type Props = {
    show: ShowDto;
}

export const ShowCard: React.FC<Props> = ({show}) => {
    return (
        <>
            <Card bordered title={<span className="font-semibold">{show.name}</span>}>
                <div className="block">
                    <div className="mb-4">{show.overview}</div>
                    <div className="w-3/4 mx-auto">
                        <Image
                            src={show.poster}
                            preview={false}
                        />
                    </div>
                </div>
            </Card>
        </>
    );
};
