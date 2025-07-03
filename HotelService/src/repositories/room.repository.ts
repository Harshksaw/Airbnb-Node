import RoomCategory from "../db/models/roomCategory";
import BaseRepository from "./base.repository";

export class RoomRepository extends BaseRepository<RoomCategory> {
    constructor() {
        super(RoomCategory);
    }


}