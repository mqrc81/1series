import React from 'react';
import { ActionIcon, Menu } from '@mantine/core';
import { withoutPropagating } from '../../utils';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEllipsisVertical, faEye } from '@fortawesome/free-solid-svg-icons';
import { useCreateTrackedShowMutation } from '../../../api';
import { DefaultProps } from '@mantine/styles/lib/theme/types/DefaultProps';

type Props = {
    showId: number;
    className: DefaultProps['className'];
}

export const ShowCardActionsDropdown: React.FC<Props> = ({showId, className}) => {
    const {mutate: createTrackedShow} = useCreateTrackedShowMutation(showId);
    return (
        <Menu
            position="bottom" closeOnItemClick offset={-3}
            classNames={{item: 'h-2 p-2', label: 'text-xs', dropdown: 'bg-dark-600/90'}}
        >
            <Menu.Target>
                <ActionIcon className={className} onClick={withoutPropagating()}>
                    <FontAwesomeIcon icon={faEllipsisVertical} />
                </ActionIcon>
            </Menu.Target>
            <Menu.Dropdown>
                <Menu.Item
                    icon={<FontAwesomeIcon icon={faEye} />}
                    onClick={withoutPropagating(() => createTrackedShow(10))}
                >
                    Track Series
                </Menu.Item>
            </Menu.Dropdown>
        </Menu>
    );
};
