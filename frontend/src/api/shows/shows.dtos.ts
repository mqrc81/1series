export type GenreDto = {
    id: number;
    genre: string;
}

export type ShowDto = {
    id: number;
    name: string;
    overview: string;
    poster: string;
    backdrop: string;
    rating: number;
    year: number;
    genres: GenreDto[];
}

export type SeasonDto = {
    showId: number;
    number: number;
    name: string;
    overview: string;
    poster: string;
    episodesCount: number;
}

export type ReleaseDto = {
    show: ShowDto;
    season: SeasonDto;
    airDate: Date;
    anticipationLevel: AnticipationLevel;
}

export enum AnticipationLevel {
    Zero,
    Moderate,
    High,
    Extreme,
}

export type ShowSearchResultDto = Pick<ShowDto, 'id' | 'name' | 'overview' | 'poster' | 'rating'>
