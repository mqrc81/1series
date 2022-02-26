import {GenreDto} from "./genre-dto";
import {NetworkDto} from "./network-dto";

export interface ShowDto {
    id?: number;
    name?: string;
    overview?: string;
    year?: number;
    poster?: string;
    rating?: number;
    homepage?: string;
    seasonsCount?: number;
    episodesCount?: number;
    genres?: GenreDto[];
    networks?: NetworkDto[];
}