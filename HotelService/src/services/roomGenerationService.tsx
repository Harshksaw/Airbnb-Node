import Room from '../db/models/room';
import { RoomGenerationJob } from '../dto/roomGeneration.dto';

import { roomCategoryRepository } from '../repositories/roomCategory.repository';

export async function generateRooms(jobData : RoomGenerationJob) {

    // Implementation for room generation\

    const { roomCategoryId, startDate, endDate, priceOverride, batchSize } = jobData;

    //check if room category exists

    const roomCategory = await roomCategoryRepository.findById(roomCategoryId)

    if(!roomCategory) {
        throw new Error(`Room category with id ${roomCategoryId} not found`);
    }
    
    const start = new Date(startDate);
    const end = new Date(endDate);
    if (start > end) {
        throw new Error('Start date must be before end date');
    }
    
}
