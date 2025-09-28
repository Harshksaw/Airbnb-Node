import { GetAvaialableRoomsDTO } from '../dto/room.dto';
import { RoomRepository } from '../repositories/room.repository';



const roomRepository = new RoomRepository();


export async function getAvailableRoomsService(getAvailableRoomsDTO: GetAvaialableRoomsDTO) {
    return await roomRepository.findByRoomCategoryIdAndDateRange(
        getAvailableRoomsDTO.roomCategoryId, new Date(getAvailableRoomsDTO.checkInDate), new Date(getAvailableRoomsDTO.checkOutDate));
}

export async function updateBookingIdToRoomsService(updateBookingIdToRoomsDTO: { bookingId: number, roomIds: number[] }) {
    return await roomRepository.updateBookingIdToRooms(updateBookingIdToRoomsDTO.roomIds, updateBookingIdToRoomsDTO.bookingId);
}