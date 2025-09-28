import { StatusCodes } from "http-status-codes";
import { getAvailableRoomsService, updateBookingIdToRoomsService } from "../services/room.service";
import { NextFunction, Request, Response } from "express";
import { GetAvaialableRoomsDTO } from "../dto/room.dto";

export async function getAvailableRoomHandler(req: Request, res: Response) {


    const rooms = await getAvailableRoomsService(req.body as unknown as GetAvaialableRoomsDTO)

    res.status(StatusCodes.OK).json({
        message:"Rooms found successfully",
        data: rooms,
        success:true
    })
    
}

export async function updateBookingIdToRoomsHandler(req: Request, res: Response, next: NextFunction) {

    const response = await updateBookingIdToRoomsService(req.body)

    res.status(StatusCodes.OK).json({
        message:"Rooms updated successfully",
        data: response,
        success:true
    })
}