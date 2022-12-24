import React from 'react';
import { Card, Image } from 'antd';
import { ReleaseDto } from '../../api';

type Props = {
    release: ReleaseDto;
}

export const ReleaseCard: React.FC<Props> = ({release}) => {
    return (
        <>
            <Card bordered title={release.show.name + ' (Season ' + release.season.number + ')'}>
                <div className="text-primary block">
                    <div className="mb-4">{release.season.overview}</div>
                    <div className="w-3/4 mx-auto">
                        <Image src={release.season.poster ?? release.show.poster} preview={false}/>
                    </div>
                </div>
            </Card>
        </>
    );
};
