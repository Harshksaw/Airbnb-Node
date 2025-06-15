import {Job, Worker} from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";


const emailProcessor = new Worker<NotificationDto>(
MAILER_QUEUE,
async(job:Job)=>{

    if(job.name !== MAILER_PAYLOAD) {
        throw new Error(`Invalid job type: ${job.name}`);
    }
    //call the service layer form here.

    

}, {
    connection: getRedisConnObject(),
    concurrency: 5, // Adjust concurrency as needed
    autorun: true, // Automatically start processing jobs
    lockDuration: 30000, // Lock duration for processing jobs
}

)


emailProcessor.on('completed', (job) => {
    console.log(`Job ${job.id} completed successfully`);
});

emailProcessor.on('failed', (job, err) => {
    console.error(`Job ${job.id} failed with error: ${err.message}`);
});
