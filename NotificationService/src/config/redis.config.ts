import Redis from "ioredis";
import { serverConfig } from ".";





export function connectToRedis() {

    try {
        const redisConfig = {
            port: serverConfig.REDIS_PORT || 6379,
            host: serverConfig.REDIS_HOST || 'localhost',
        }
        const connection = new Redis(redisConfig);
        return connection;

    } catch (error) {
        console.error("Error connecting to Redis:", error);
        
    }

}