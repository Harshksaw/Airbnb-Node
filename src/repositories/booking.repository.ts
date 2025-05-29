import { Prisma } from '@prisma/client';
import  prismaClient  from '../prisma/client';
import { generateIdempotencyKey } from '../utils/generateIdempotencyKey';



export async function createBooking(bookingData: Prisma.BookingCreateInput) {
        const booking = await prismaClient.booking.create({
            data: bookingData,
        });
    
        // logger.info(`Booking created: ${booking.id}`);
    
        return booking;
    }


export async function createIdempotencyKey(bookingId: number, key: number) {
        const idempotencyKey = await prismaClient.idempotencyKey.create({
            data: {
                key,
                bookingId:{
                    connect:{
                        id:bookingId
                    }
                }
                
            },
        });
    
        // logger.info(`Idempotency key created for booking ${bookingId}: ${idempotency.key}`);
    
        return idempotencyKey;
    }



export async function getIdempotencyKey(key:string){

    const idempotencyKey = await prismaClient.idempotencyKey.findUnique({
        where:{
            key
        }
    })
    return idempotencyKey
}

// export async function changeBookingStatus(bookingId: number, status:Prisma.EnumBookingStatusFieldUpdateOperationsInput){

//     const booking= await prismaClient.booking.update({
//         where:{
//             id:bookingId
//         },
//         data:{
//             status:status
//         }
//     })
//     return booking
// }

export async function confirmBooking(bookingId :number) {
    const booking= await prismaClient.booking.update({
                where:{
                    id:bookingId
                },
                data:{
                    status:"CANCELLED"
                }
            })
            return booking
    
}


export async function finalizedIdempotencyKey(key:string){
    const idempotencyKey= await prismaClient.idempotencyKey.update({
        where:{
            key
        },
        data:{
            finalized:true
        }
    })
    return idempotencyKey
}