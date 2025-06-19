import {Job, Worker} from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
// import { MAILER_PAYLOAD } from "../producers/email.producer";
import { renderMailTemplate } from "../templates/templates.handler";
import logger from "../config/logger.config";
import { sendEmail } from "../services/mailer.service";




export const setupMailerWorker = () =>{

const emailProcessor = new Worker<NotificationDto>(
MAILER_QUEUE,
async(job:Job)=>{

    // if(job.name !== MAILER_PAYLOAD) {
    //     throw new Error(`Invalid job type: ${job.name}`);
    // }
    //call the service layer form here.
    const payload = job.data;
    console.log(`Processing email job for ${JSON.stringify(payload)}`);

    const emailContent = await renderMailTemplate(payload.templateId, payload.params);

    await sendEmail(payload.to, payload.subject, emailContent);

    logger.info(`Email sent successfully to ${payload.to} with subject: ${payload.subject}`);


}, {
    connection: getRedisConnObject(),

}

)


emailProcessor.on('completed', (job) => {
    console.log(`Job ${job.id} completed successfully`);
});

emailProcessor.on('failed', (job, err) => {
    console.error(`Job ${job} failed with error: ${err.message}`);
});


}
