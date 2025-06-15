

import Redis from "ioredis";
import { serverConfig } from ".";

export function connectToRedis() {
    try {
        let connection: Redis;

        const redisConfig = {
            port: serverConfig.REDIS_PORT || 6379,
            host: serverConfig.REDIS_HOST || 'localhost',
        }

        return () => {
            if (!connection) {
                connection = new Redis(redisConfig);
                console.log("Connected to Redis:", redisConfig);
            }
            return connection;
        }

    } catch (error) {
        console.error("Error connecting to Redis:", error);
        throw error; // Re-throw error to handle it at caller level
    }
}

export const getRedisConnObject = connectToRedis();
