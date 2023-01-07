import { useFiltersStore } from '../../stores';
import { Genre } from '../../api';

export const useGenresFilter = () => {
    const filteredGenres = useFiltersStore(state => state.filteredGenres);

    return {
        filterGenre: (genre: Genre) => filteredGenres.push(genre.name),
        unfilterGenre: (genre: Genre) => filteredGenres.splice(filteredGenres.indexOf(genre.name), 1),
        isGenreFiltered: (genre: Genre) => filteredGenres.includes(genre.name),
        isGenreFilterActive: () => filteredGenres.length > 0,
    };
};
