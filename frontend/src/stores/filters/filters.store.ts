import create from 'zustand';

type FiltersState = {
    filteredGenres: string[];
}

export const useFiltersStore = create<FiltersState>((set) => ({
    filteredGenres: [],
} as FiltersState));

