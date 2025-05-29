import {Request, Response} from 'express'

export const createBookingHandler = async (req:Request, res:Response){
 
    const booking = await createBookingHandler(
        req.body.userId, req.body.hotelId, req.body.totalGuests, req.body.bookingAmount
    )
}