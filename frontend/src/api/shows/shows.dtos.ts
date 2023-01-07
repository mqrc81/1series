export type Genre = {
    id: number;
    name: string;
}

export type Show = {
    id: number;
    name: string;
    overview: string;
    poster: string;
    backdrop: string;
    rating: number;
    year: number;
    genres: Genre[];
}

export type Season = {
    showId: number;
    number: number;
    name: string;
    overview: string;
    poster: string;
    episodesCount: number;
}

export type ReleaseDto = {
    show: Show;
    season: Season;
    airDate: Date;
    anticipationLevel: AnticipationLevel;
}

export enum AnticipationLevel {
    Zero,
    Moderate,
    High,
    Extreme,
}

export type ShowSearchResult = Pick<Show, 'id' | 'name' | 'overview' | 'poster' | 'backdrop' | 'rating'>
