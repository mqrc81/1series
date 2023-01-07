import create from 'zustand';
import { User } from '../../api';
import { persist } from 'zustand/middleware';

type AuthState = {
    user: User;
    isLoggedIn: () => boolean
    login: (user: User) => void;
    logout: () => void;
}

export const useAuthStore = create(persist<AuthState>(
    (set, get) => ({
        isLoggedIn: () => !!get().user?.id,
        login: (user: User) => set({user}),
        logout: () => set({user: null}),
    } as AuthState),
    {
        name: 'auth',
        getStorage: () => sessionStorage,
    },
));
