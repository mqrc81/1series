export type User = {
    id: number;
    username: string;
    email: string;
    password: string;
    emailVerified: boolean;
} & NotificationOptions

export type NotificationOptions = {
    notifyReleases: boolean;
    notifyRecommendations: boolean;
}

export type FailedImdbImport = {
    imdbId: string;
    title: string;
    reason: string;
}

export type SignUserInDto = Pick<User, 'email' | 'password'>

export type SignUserUpDto = Pick<User, 'username' | 'email' | 'password'>

export type TrackShowDto = {
    showId: number;
    rating?: number;
}
