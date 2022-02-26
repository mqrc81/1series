import {Network} from "./network";
import {ShowDto} from "../dtos/show-dto";
import {Genre} from "./genre";

export interface Show extends ShowDto {
    genres?: Genre[];
    networks?: Network[];
}