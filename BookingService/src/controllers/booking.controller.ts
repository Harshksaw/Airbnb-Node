import {Request, Response} from 'express'
import { createBookingService } from '../services/booking.service'

export const createBookingHandler = async (req:Request, res:Response){
 
    const booking = await createBookinService(req.body);

    res.status(201).json({
        bookingId: booking.bookingId,
        idempotencyKey: booking.idempotencyKey,
    })
}

export const confirmBookingHandler = async (req:Request, res:Response){
 
    const booking = await createBookingService(req.params.idempotencyKey)
    res.status(201).json({
        bookingId: booking.id,
        idempotencyKey: booking.status,
    })
}