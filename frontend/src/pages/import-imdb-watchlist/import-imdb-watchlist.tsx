import React from 'react';
import { useToast } from '../../hooks/use-toast/use-toast';
import { useImportImdbWatchlistMutation } from '../../api';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileImport } from '@fortawesome/free-solid-svg-icons';

// TODO ms: this component is only temporary, move to user dropdown with modal or something later
const ImportImdbWatchlist: React.FC = () => {
    const {successToast, errorToast, toastContextHolder} = useToast();
    const {
        data: failedImports = [],
        mutate: uploadFile,
        isSuccess, isLoading,
    } = useImportImdbWatchlistMutation({
        onSuccess: (response) =>
            successToast('Successfully added your IMDb watchlist to tracked shows (' + response.length + ' failed)'),
        onError: () => errorToast('Error importing IMDb watchlist...'),
    });

    return (
        <>
            <div className="grid grid-cols-1 w-full">
                <label>
                    <input
                        type="file"
                        onChange={({target}) => uploadFile(target.files[0])}
                        disabled={isSuccess || isLoading}
                    />
                    <FontAwesomeIcon icon={faFileImport} size="6x"/>
                    <p className="ant-upload-hint">
                        Upload your WATCHLIST.csv file
                    </p>
                </label>
                {isSuccess && failedImports.length > 0 && <div>
                    <span>Failed Imports: </span>
                    {failedImports.map(({title, reason}) => (
                        <div>{title}: {reason}</div>
                    ))}
                </div>}
            </div>
            {toastContextHolder}
        </>
    );
};

export default ImportImdbWatchlist;
