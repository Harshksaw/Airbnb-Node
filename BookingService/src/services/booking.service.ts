
import { confirmBooking, createBooking, createIdempotencyKey, finalizedIdempotencyKey, getIdempotencyKey } from '../repositories/booking.repository';
import { generateIdempotencyKey } from '../utils/generateIdempotencyKey';
import { BadRequestError, NotFoundError } from '../utils/errors/app.error';
import { CreateBookingDTO } from '../dto/booking.dto';



export async function createBookingService(
    createBookingDto: CreateBookingDTO

) {
    const booking = await createBooking({
        userId: createBookingDto.userId,
        hotelId: createBookingDto.hotelId,
        totalGuests: createBookingDto.totalGuests,
        bookingAmount: createBookingDto.bookingAmount

    })
    const idempotencyKey = generateIdempotencyKey();

    await createIdempotencyKey(idempotencyKey, booking.id)

    return {
        bookingId: booking.id,
        idempotencyKey: idempotencyKey
    }



}



export async function finalizedBookingService(idempotencyKey: string) {

    const idempotencyKeyData = await getIdempotencyKey(idempotencyKey)

    if (!idempotencyKeyData) {
        throw new NotFoundError('Idempotency key not found')

    }
    if (idempotencyKeyData.finalized) {
        throw new BadRequestError('Idempotency key already finalized');
    }

    const booking = await confirmBooking(idempotencyKeyData.bookingId);

    await finalizedIdempotencyKey(idempotencyKey)

    return booking;
}