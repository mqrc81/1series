export type User = {
    id: number;
    username: string;
    email: string;
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

export type LoginUserDto = {
    emailOrUsername: string;
    password: string;
}

export type RegisterUserDto = Pick<User, 'username' | 'email'> & {
    password: string;
}

export type TrackShowDto = {
    showId: number;
    rating?: number;
}
