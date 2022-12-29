import React, { useState } from 'react';
import { useToast } from '../../hooks/use-toast/use-toast';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileImport } from '@fortawesome/free-solid-svg-icons';
import { Spin, Upload } from 'antd';
import { UploadChangeParam, UploadFile } from 'antd/es/upload/interface';
import { FailedImdbImport } from '../../api';

// TODO ms: this component is only temporary, move to user dropdown with modal or something later
const ImportImdbWatchlist: React.FC = () => {
    const {successToast, errorToast, toastContextHolder} = useToast();
    const [isLoading, setIsLoading] = useState(false);
    const [failedImports, setFailedImports] = useState<FailedImdbImport[]>([]);

    return (
        <>
            <Upload.Dragger
                action={import.meta.env.VITE_BACKEND_URL + '/api/users/importImdbWatchlist'}
                withCredentials={true}
                name="file"
                accept=".csv"
                disabled={isLoading}
                onChange={({file: {status, response: failedImports}}: UploadChangeParam<UploadFile<FailedImdbImport[]>>) => {
                    setIsLoading(status === 'uploading');
                    switch (status) {
                        case 'done':
                            setFailedImports(failedImports);
                            return successToast('Successfully added your IMDb watchlist to tracked shows (' + (failedImports.length) + ' failed)');
                        case 'error':
                            return errorToast('Error importing IMDb watchlist...');
                    }
                }}
            >
                {isLoading
                    ? <Spin className="m-auto mt-5" spinning size="large"/>
                    : <>
                        <p className="ant-upload-drag-icon">
                            <FontAwesomeIcon icon={faFileImport} size="4x"/>
                        </p>
                        <p className="ant-upload-text">Click or drag file to this area to upload</p>
                        <p className="ant-upload-hint">Only the exported WATCHLIST.csv is supported</p>
                    </>}

            </Upload.Dragger>
            {!isLoading && failedImports.length > 0 && <div>
                <span>Failed Imports: </span>
                {failedImports.map(({title, reason}) => (
                    <div>{title}: {reason}</div>
                ))}
            </div>}
            {toastContextHolder}
        </>
    );
};

export default ImportImdbWatchlist;
