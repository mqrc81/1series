import React from 'react';
import { useToast } from '../../hooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useImportImdbWatchlistMutation } from '../../api';
import { Group } from '@mantine/core';
import { Dropzone, MIME_TYPES } from '@mantine/dropzone';
import { faFileImport } from '@fortawesome/free-solid-svg-icons';
import { faCircleCheck, faCircleXmark } from '@fortawesome/free-regular-svg-icons';

const ONE_MEGABYTE = 1024 ** 2;

// TODO ms: this component is only temporary, move to user dropdown with modal or something later
const ImportImdbWatchlist: React.FC = () => {
    const {successToast, errorToast} = useToast();
    const {
        data: failedImports = [],
        mutate: uploadFile,
        isSuccess, isLoading,
    } = useImportImdbWatchlistMutation({
        onSuccess: (failedImports) => successToast('Successfully added your IMDb watchlist to tracked shows (' + (failedImports.length) + ' failed)'),
    });

    return (
        <>
            <Dropzone
                name="file"
                accept={[MIME_TYPES.csv]}
                maxSize={ONE_MEGABYTE}
                disabled={isLoading}
                onDrop={(files) => uploadFile(files[0])}
                onReject={() => errorToast('Please select a valid .csv file')}
            >
                <Group position="center" spacing="xl" className="h-52">
                    <Dropzone.Accept>
                        <FontAwesomeIcon className="text-green-600" icon={faCircleCheck} size="4x"/>
                    </Dropzone.Accept>
                    <Dropzone.Reject>
                        <FontAwesomeIcon className="text-red-600" icon={faCircleXmark} size="4x"/>
                    </Dropzone.Reject>
                    <Dropzone.Idle>
                        <FontAwesomeIcon className="text-violet-600" icon={faFileImport} size="4x" />
                    </Dropzone.Idle>
                    <div>
                        <div className="text-xl">
                            Export your IMDb watchlist and drag the file here or click to select
                        </div>
                        <div className="text-sm text-gray-500 mt-3">
                            The file must not exceed 1mb
                        </div>
                    </div>
                </Group>
            </Dropzone>
            {isSuccess && failedImports.length > 0 && <div>
                <span>Failed Imports: </span>
                {failedImports.map(({title, reason}) => (
                    <div>{title}: {reason}</div>
                ))}
            </div>}
        </>
    );
};

export default ImportImdbWatchlist;
