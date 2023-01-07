import { useFiltersStore } from '../../stores';
import { GenreDto } from '../../api';

export const useGenresFilter = () => {
    const filteredGenres = useFiltersStore(state => state.filteredGenres);

    return {
        filterGenre: (genre: GenreDto) => filteredGenres.push(genre.name),
        unfilterGenre: (genre: GenreDto) => filteredGenres.splice(filteredGenres.indexOf(genre.name), 1),
        isGenreFiltered: (genre: GenreDto) => filteredGenres.includes(genre.name),
        isGenreFilterActive: () => filteredGenres.length > 0,
    };
};
