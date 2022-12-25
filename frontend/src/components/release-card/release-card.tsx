import React from 'react';
import { Card, Image } from 'antd';
import { ReleaseDto } from '../../api';

type Props = {
    release: ReleaseDto;
}

export const ReleaseCard: React.FC<Props> = ({release}) => {
    return (
        <>
            <Card bordered title={<span className="font-semibold">{release.show.name + ' (' + release.season.name + ')'}</span>}>
                <div className="block">
                    <div className="mb-4">{release.season.overview ?? 'No season overview...'}</div>
                    <div className="w-3/4 mx-auto">
                        <Image
                            src={release.season.poster ?? release.show.poster}
                            preview={false}
                        />
                    </div>
                </div>
            </Card>
        </>
    );
};
