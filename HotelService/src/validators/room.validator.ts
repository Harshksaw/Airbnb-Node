import { z} from "zod";



export const GetAvaialableRoomsSchema = z.object({
    roomCategory: z.number({message: "Room category ID must be a number"}),
    checkInDate: z.string({message: "Check-in date is required"}),
    checkOutDate: z.string({message: "Check-out date is required"}),
})

export const updateBookingIdToRoomsSchema = z.object({
    roomIds: z.array(z.number({message: "Room ID must be a number"}), {required_error: "Room IDs are required"}),
    bookingId: z.number({message: "Booking ID must be a number"})
})