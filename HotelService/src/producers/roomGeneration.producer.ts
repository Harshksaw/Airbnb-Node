import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { roomGenerationQueue } from '../queues/roomGeneration.queue';


export const ROOM_GENERATION_PAYLOAD = 'room-generation-payload';

export const addRoomGenerationJobToQueue = async(payload:RoomGenerationJob)=>{
    await roomGenerationQueue.add(ROOM_GENERATION_PAYLOAD,payload);


}
