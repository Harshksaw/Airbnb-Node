import {z} from "zod";


export const RoomGenerationDTO = z.object({

    roomCategoryId: z.string().uuid(),
    startDate: z.string().datetime(),
    endDate: z.string().datetime(),
    scheduleType: z.enum(['immidiate', 'weekly', 'monthly']).default('immidiate'),
    scheduledAt: z.string().datetime().optional(),

    priceOverride: z.number().positive().optional(),


})


export const RoomGenerationJob = z.object({
    roomCategoryId: z.number().positive(),
    startDate: z.string().datetime(),
    endDate: z.string().datetime(),
    priceOverride: z.number().positive().optional(),
    batchSize: z.number().positive().default(100),
})

export interface RoomGenerationResponse{
    success: boolean;
    totalRoomsCreated: number;
    totalDatesProcessed: number;
    errors: string[];
    jobId: string;
}

export type RoomGenerationJob = z.infer<typeof RoomGenerationJob>;

export type RoomGenerationRequest = z.infer<typeof RoomGenerationDTO>;
export type RoomGenerationJobRequest = z.infer<typeof RoomGenerationJob>;