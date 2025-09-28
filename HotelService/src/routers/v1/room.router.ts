import express from 'express';
import { getAvailableRoomHandler, updateBookingIdToRoomsHandler } from '../../controllers/room.controller';
import {  validateQueryParams, validateRequestBody } from '../../validators';
import { GetAvaialableRoomsSchema, updateBookingIdToRoomsSchema } from '../../validators/room.validator';

const roomRouter = express.Router();

roomRouter.get('/available',validateQueryParams(GetAvaialableRoomsSchema), getAvailableRoomHandler); 
roomRouter.post('/update-booking',validateRequestBody(updateBookingIdToRoomsSchema), updateBookingIdToRoomsHandler);


export default roomRouter;