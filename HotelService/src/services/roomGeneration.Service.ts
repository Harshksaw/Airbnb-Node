import Room from '../db/models/room';
import { RoomGenerationJob } from '../dto/roomGeneration.dto';

import { RoomCategoryRepository } from '../repositories/roomCategory.repository';
import RoomCategory from '../db/models/roomCategory';
import { RoomRepository } from '../repositories/room.repository';

const roomRepostitory = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();

export async function generateRooms(jobData : RoomGenerationJob) {

    // Implementation for room generation\

    const { roomCategoryId,  priceOverride, batchSize } = jobData;

    //check if room category exists

    const roomCategory = await roomCategoryRepository.findById(roomCategoryId)

    if(!roomCategory) {
        throw new Error(`Room category with id ${roomCategoryId} not found`);
    }
    
    const startDate = new Date(jobData.startDate);
    const endDate = new Date(jobData.endDate);
    if (startDate > endDate) {
        throw new Error('Start date must be before end date');
    }

    if (startDate < new Date()){
        throw new Error('Start date must be in the future');
    }

    const totalDays =  Math.ceil(endDate.getTime() - startDate.getTime()) /(1000*60*60*24)

}



export async function processDateBatch(  roomCategory : RoomCategory ,priceOverride: number | undefined, startDate: Date, endDate : Date ,batchSize: number) {
    
    
    let roomsCreated = 0;
    let datesProcessed = 0;
    const roomsToCreate:any[] = []

    const currentDate = new Date(startDate);

    while(currentDate <= endDate ){
        const exisitingRooms = await roomRepostitory.findByRoomCategoryIdAndDate(
            roomCategory.id,
            currentDate
        )
        if(!exisitingRooms) {
            roomsToCreate.push({
                hotelId: roomCategory.hotelId,
                roomCategoryId: roomCategory.id,
                dateOfAvailability: new Date(currentDate),
                price: priceOverride || roomCategory.price,

            })
        }
    }
    currentDate.setDate(currentDate.getDate() + 1);
    datesProcessed++;

}