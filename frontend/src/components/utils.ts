import { MouseEvent } from 'react';

export const withoutPropagating = (func = () => {}) => (e: MouseEvent) => {
    e.stopPropagation();
    func();
};
