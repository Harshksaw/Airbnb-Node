
import { createRoomCategoryDTO } from '../dto/roomCategory.dto';

import { RoomCategoryRepository } from '../repositories/roomCategory.repository';
import { HotelRepository } from '../repositories/hotel.repository';

const roomCategoryRepository = new RoomCategoryRepository();
const hotelRepository = new HotelRepository();

export async function createRoomCategoryService(createRoomCategoryDTO: createRoomCategoryDTO) {
    const RoomCategory = await roomCategoryRepository.create(createRoomCategoryDTO);
    return RoomCategory;
}

export async function getRoomCategoryByIdService(id: number) {
    const roomCategory = await roomCategoryRepository.findById(id);
    if (!roomCategory) {
        throw new Error(`Room category with id ${id} not found`);
    }
    return roomCategory;
}

export async function getAllRoomCategoriesByHotelIdService(hotelId:number) {
    const hotel = await hotelRepository.findById(hotelId);

    if(!hotel){
        throw new Error(`Hotel with id ${hotelId} not found`);
    }
    const roomCategories = await roomCategoryRepository.findAllByHotelId(hotelId);

    return roomCategories;

}
    

export async function deleteRoomCategoryService(id: number) {
    const roomCategory = await roomCategoryRepository.findById(id);
    if (!roomCategory) {
        throw new Error(`Room category with id ${id} not found`);
    }
    await roomCategoryRepository.delete({ id });
    return { message: `Room category with id ${id} deleted successfully` };
}
