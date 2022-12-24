export type Paginated<T> = T & {
    nextPage: number;
    previousPage?: number;
}
