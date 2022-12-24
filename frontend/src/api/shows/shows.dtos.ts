export type ShowDto = {
    id: number;
    name: string;
    overview: string;
    poster: string;
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
