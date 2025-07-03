
import RoomCategory from "../db/models/roomCategory";
import { createRoomCategoryDTO } from "../dto/roomCategory.dto";
import BaseRepository from "./base.repository";


export class RoomCategoryRepository  extends BaseRepository<RoomCategory>{

    constructor() {
        super(RoomCategory);
    }

    async create(data: createRoomCategoryDTO): Promise<RoomCategory> {

        const exists = await this.model.findOne({
            where: {
                hotelId: data.hotelId,
                roomType: data.roomType
            }
        })
        if(!exists) {
            throw new Error(`Room category with hotelId ${data.hotelId} and roomType ${data.roomType} already exists`);
        }

        const roomCategory = await this.model.create(data);
        
        return roomCategory;
    }

    async findAllByHotelId(hotelId: number): Promise<RoomCategory[]> {
        const roomCategories = await this.model.findAll({
            where: {
                hotelId: hotelId,
                deletedAt: null // Assuming soft delete is implemented with a deletedAt field
            }
        });

        if (!roomCategories || roomCategories.length === 0) {
            throw new Error(`No room categories found for hotel with id ${hotelId}`);
        }

        return roomCategories;
    }

}

