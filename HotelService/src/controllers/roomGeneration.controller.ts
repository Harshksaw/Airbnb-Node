import { Request, Response, NextFunction } from "express";

import { StatusCodes } from "http-status-codes";
import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer";



export async function generateRoomHandler(req: Request, res: Response) {


    const result = await addRoomGenerationJobToQueue(req.body);

    res.status(StatusCodes.OK).json({
        message: "Rooms generated successfully",
        success: true,
        data: result,
    })
}