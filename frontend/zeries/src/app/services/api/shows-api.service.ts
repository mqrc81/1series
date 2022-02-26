import {Injectable} from '@angular/core';
import {Show} from "../../interfaces/domain/show";

@Injectable({
    providedIn: 'root'
})
export class ShowsApiService {

    constructor() {
    }

    getShow(showId: number): Promise<Show> {
        console.log(showId)
        return new Promise<Show>(resolve => {
            setTimeout(
                () => resolve({
                    id: 3,
                    name: 'Abc',
                }),
                300
            );
        });
    }
}
