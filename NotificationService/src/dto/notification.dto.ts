export interface NotificationDto {
    to: string;
    subject:string;
    templatedId: string:
    params: Record<string, any>;
}

