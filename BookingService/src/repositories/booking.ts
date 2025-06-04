import {Booking, BookingStatus} from '@PrismaClient';
import  prismaClient, { Prisma }  from '@prisma/client';



export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
  try {
    const booking = await prismaClient.booking.create({
      data: bookingInput,
    });
    return booking;
  } catch (error) {
    console.error('Error creating booking:', error);
    throw error;
  }
}