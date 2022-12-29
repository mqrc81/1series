import React from 'react';
import { useToast } from '../../hooks/use-toast/use-toast';
import { useImportImdbWatchlistMutation } from '../../api';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileImport } from '@fortawesome/free-solid-svg-icons';
import { Spin, Upload } from 'antd';

// TODO ms: this component is only temporary, move to user dropdown with modal or something later
const ImportImdbWatchlist: React.FC = () => {
    const {successToast, errorToast, toastContextHolder} = useToast();
    const {
        data: failedImports = [],
        mutate: uploadFile,
        isSuccess, isLoading,
    } = useImportImdbWatchlistMutation({
        onSuccess: (response) => successToast('Successfully added your IMDb watchlist to tracked shows (' + response.length + ' failed)'),
        onError: () => errorToast('Error importing IMDb watchlist...'),
    });

    return (
        <>
            <Upload.Dragger
                beforeUpload={(file) => {
                    uploadFile(file);
                    return false;
                }}
                accept=".csv"
                disabled={isLoading}
            >
                {isLoading
                    ? <Spin className="m-auto mt-5" spinning size="large"/>
                    : <FontAwesomeIcon icon={faFileImport} size="6x"/>}

            </Upload.Dragger>
            {isSuccess && failedImports.length > 0 && <div>
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
