import { Request, Response, NextFunction } from "express";

import { StatusCodes } from "http-status-codes";
import { generateRooms } from "../services/roomGeneration.Service";

export async function generateRoomHandler(req: Request, res: Response) {


    const result = await generateRooms(req.body);

    res.status(StatusCodes.OK).json({
        message: "Rooms generated successfully",
        success: true,
        data: result,
    })
}