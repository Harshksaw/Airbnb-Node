import { prismaClient } from '@prisma/client';
import { createBooking } from '../repositories/booking.repository';



export async function createBookinService(
userId:number,
hotelId:number,
totalGuests:number,
bookingAmount:number

) {
    const booking = await createBooking({
        userId,
        hotelId,
        totalGuests:totalGuests,
        bookingAmount:bookingAmount

    })


    
}



export async function finalizedBookingService(){
    
}