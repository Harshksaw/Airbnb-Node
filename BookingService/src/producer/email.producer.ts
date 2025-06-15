import { NotificationDto } from "../dto/notification.dto";
import { mailerQueue } from "../queues/email.queue";




export const MAILER_PAYLOAD = "payload:mail";
export const addEmailToQueue = async(payload:NotificationDto) => {
    try {
        await mailerQueue.add('sendEmail', payload, {
            attempts: 3, // Retry up to 3 times on failure
            backoff: {
                type: 'exponential',
                delay: 1000, // Initial delay of 1 second
            },
        });
        console.log(`Email job added to queue for ${payload.to}`);
    } catch (error) {
        console.error('Error adding email job to queue:', error);
    }
}